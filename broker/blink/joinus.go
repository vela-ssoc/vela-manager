package blink

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dfcfw/spdy"
	"github.com/vela-ssoc/manager/infra/conf"
	"github.com/vela-ssoc/manager/inward/evtrsk"
	"github.com/vela-ssoc/manager/libkit/httpclient"
	"github.com/vela-ssoc/manager/model"
	"gorm.io/gorm"
)

var (
	ErrBrokerNotFound    = errors.New("broker 节点不存在")
	ErrBrokerRepeat      = errors.New("broker 节点重复连接")
	ErrBrokerInet        = errors.New("broker IP 不合法")
	ErrBrokerUnconnected = errors.New("broker 节点尚未连接")
)

type Huber interface {
	Joiner
	Reset()
	Fetch(context.Context, int64, Operator, io.Reader) (*http.Response, error)
	Through(bid string, op Operator, req *http.Request, res http.ResponseWriter) error
}

// Hub broker 节点的连接中心
func Hub(db *gorm.DB, notice evtrsk.Handler, handler http.Handler, cfg conf.Config) Huber {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	hub := &brkHub{
		db:      db,
		notice:  notice,
		handler: handler,
		brokers: make(map[int64]*connect, 16), // 一般不会超过 16 个 broker
		cfg:     cfg,
		random:  random,
	}
	transport := &http.Transport{DialContext: hub.dialContext}
	cli := &http.Client{Transport: transport}
	client := httpclient.NewClient(cli)
	hub.client = client

	return hub
}

type brkHub struct {
	db      *gorm.DB
	notice  evtrsk.Handler
	handler http.Handler
	mutex   sync.RWMutex
	brokers map[int64]*connect
	client  httpclient.Client
	cfg     conf.Config
	random  *rand.Rand
}

// Auth 鉴权授权管理
func (bh *brkHub) Auth(ident Ident) (any, http.Header, error) {
	id, secret, inet := ident.ID, ident.Secret, ident.IP
	if len(inet) == 0 || inet.IsLoopback() {
		return nil, nil, ErrBrokerInet
	}

	// 查询 broker
	var brk model.Broker
	ipv4 := inet.String()
	if err := bh.db.Take(&brk, "id = ? AND inet = ? AND secret = ?", id, ipv4, secret).
		Error; err != nil {
		return nil, nil, ErrBrokerNotFound
	}
	if brk.Status || bh.getConn(id) != nil {
		return nil, nil, ErrBrokerRepeat
	}

	// 随机生成一个 32-64 位的加密密钥
	psz := bh.random.Intn(33) + 32
	passwd := make([]byte, psz)
	_, _ = bh.random.Read(passwd)
	issue := Issue{
		Name:     brk.Name,
		Passwd:   passwd,
		Listen:   Listen{Addr: ":8180"},
		Logger:   bh.cfg.Logger,
		Database: bh.cfg.Database,
	}

	return issue, nil, nil
}

func (bh *brkHub) Join(tran net.Conn, ident Ident, ret any) error {
	issue := ret.(Issue)
	mux := spdy.Server(tran, spdy.WithEncrypt(issue.Passwd))

	id := ident.ID
	conn := &connect{
		ident:  ident,
		issue:  issue,
		waiter: bh,
		mux:    mux,
	}

	if !bh.putConn(conn) { // [上线] 将 connect 存到连接池中
		return ErrBrokerRepeat
	}
	defer bh.delConn(id) // [下线] 删除连接池中的

	now := time.Now()
	nowAt := sql.NullTime{Valid: true, Time: now}
	semver := model.Semver(ident.Semver)
	tbl := &model.Broker{
		ID:         id,
		MAC:        ident.MAC,
		Goos:       ident.Goos,
		Arch:       ident.Arch,
		CPU:        ident.CPU,
		Semver:     semver,
		Status:     true,
		PID:        ident.PID,
		Workdir:    ident.Workdir,
		Executable: ident.Executable,
		Username:   ident.Username,
		Hostname:   ident.Hostname,
		PingAt:     nowAt,
		JoinAt:     nowAt,
	}
	if err := bh.db.UpdateColumns(tbl).Error; err != nil { // [上线] 修改为在线状态
		return err
	}
	defer func() {
		bh.db.Model(tbl).Update("status", false) // [下线] 修改为离线状态
	}()

	// 通知上线
	inet := ident.IP.String()
	online := &model.Event{
		MinionID:  id,
		Inet:      inet,
		Msg:       "代理节点上线",
		Typeof:    "broker.online",
		SendAlert: true,
		OccurAt:   now,
		CreatedAt: now,
	}
	_ = bh.notice.Event(online)

	srv := &http.Server{
		Handler: bh.handler,
		BaseContext: func(net.Listener) context.Context {
			return context.WithValue(context.Background(), brokerCtxKey, conn)
		},
	}
	_ = srv.Serve(mux) // 此处会阻塞，一旦执行结束说明连接断开

	// 通知上线
	offline := &model.Event{
		MinionID:  id,
		Inet:      inet,
		Subject:   "代理节点下线",
		Msg:       "代理节点下线",
		Typeof:    "broker.offline",
		SendAlert: true,
		OccurAt:   now,
		CreatedAt: now,
	}
	_ = bh.notice.Event(offline)

	return nil
}

func (bh *brkHub) Reset() {
	bh.db.Model(&model.Broker{}).
		Where("status = ?", true).
		UpdateColumn("status", false)
}

func (bh *brkHub) Fetch(ctx context.Context, id int64, op Operator, body io.Reader) (*http.Response, error) {
	req := bh.newRequest(ctx, id, op, body)
	return bh.client.Fetch(req)
}

// Through ss
func (bh *brkHub) Through(bid string, op Operator, req *http.Request, res http.ResponseWriter) error {
	req.URL.Scheme = "http"
	req.URL.Host = bid
	req.URL.Path = op.Path()
	req.RequestURI = ""
	ret, err := bh.client.Fetch(req)
	if err != nil {
		return err
	}
	for k, v := range ret.Header {
		res.Header().Set(k, strings.Join(v, ", "))
	}
	res.WriteHeader(ret.StatusCode)
	_, err = io.Copy(res, ret.Body)

	return err
}

// getConn 通过 ID 获取连接
func (bh *brkHub) getConn(id int64) *connect {
	bh.mutex.RLock()
	conn := bh.brokers[id]
	bh.mutex.RUnlock()

	return conn
}

// putConn 存入一个 broker 连接
// false- 连接已经存在，存入失败
func (bh *brkHub) putConn(conn *connect) bool {
	id := conn.ID()
	bh.mutex.Lock()
	_, exist := bh.brokers[id]
	if !exist {
		bh.brokers[id] = conn
	}
	bh.mutex.Unlock()

	return !exist
}

// delConn 删除 broker 连接
func (bh *brkHub) delConn(id int64) bool {
	bh.mutex.Lock()
	_, exist := bh.brokers[id]
	if exist {
		delete(bh.brokers, id)
	}
	bh.mutex.Unlock()

	return exist
}

func (bh *brkHub) dialContext(_ context.Context, _, addr string) (net.Conn, error) {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, net.InvalidAddrError(addr)
	}
	id, _ := strconv.ParseInt(host, 10, 64)
	if conn := bh.getConn(id); conn != nil {
		return conn.mux.Dial()
	}

	return nil, ErrBrokerUnconnected
}

func (*brkHub) newRequest(ctx context.Context, id int64, op Operator, body io.Reader) *http.Request {
	method := op.Method()
	path := op.Path()

	host := strconv.FormatInt(id, 10)
	addr := &url.URL{Scheme: "http", Host: host, Path: path}
	req := &http.Request{
		Method:     method,
		URL:        addr,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
	}

	if v, ok := body.(interface{ Len() int }); ok {
		req.ContentLength = int64(v.Len())
	}
	switch v := body.(type) {
	case nil:
	case io.ReadCloser:
		req.Body = v
	case *bytes.Buffer:
		req.Body = io.NopCloser(v)
	case *bytes.Reader:
		req.Body = io.NopCloser(v)
	case *strings.Reader:
		req.Body = io.NopCloser(v)
	default:
		req.ContentLength = -1
		req.Body = io.NopCloser(body)
	}

	// For client requests, Request.ContentLength of 0
	// means either actually 0, or unknown. The only way
	// to explicitly say that the ContentLength is zero is
	// to set the Body to nil. But turns out too much code
	// depends on NewRequest returning a non-nil Body,
	// so we use a well-known ReadCloser variable instead
	// and have the http package also treat that sentinel
	// variable to mean explicitly zero.
	if req.Body != nil && req.ContentLength == 0 {
		req.Body = http.NoBody
	}

	if ctx == nil {
		ctx = context.Background()
	}

	return req.WithContext(ctx)
}
