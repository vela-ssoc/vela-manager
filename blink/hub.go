package blink

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/vela-ssoc/backend-common/httpclient"
	"github.com/vela-ssoc/backend-common/model"
	"github.com/vela-ssoc/backend-common/opurl"
	"github.com/vela-ssoc/backend-common/pubrr"
	"github.com/vela-ssoc/backend-common/spdy"
	"github.com/vela-ssoc/vela-manager/infra/conf"
	"github.com/vela-ssoc/vela-manager/inward/evtrsk"
	"gorm.io/gorm"
)

var (
	ErrBrokerNotFound = errors.New("broker 节点不存在")
	ErrBrokerRepeat   = errors.New("broker 节点重复连接")
	ErrBrokerInet     = errors.New("broker IP 不合法")
	ErrBrokerOffline  = errors.New("代理节点未上线")
)

type Huber interface {
	Joiner
	ResetDB() error
	CallB(context.Context, opurl.URLer, io.Reader) (*http.Response, error)
	JSONB(context.Context, opurl.URLer, any, any) error
	OnewayB(context.Context, opurl.URLer, io.Reader) error
	Forward(opurl.URLer, http.ResponseWriter, *http.Request)
}

// Hub broker 节点的连接中心
func Hub(db *gorm.DB, notice evtrsk.Handler, handler http.Handler, cfg conf.Config) Huber {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	hub := &brkHub{
		db:      db,
		notice:  notice,
		handler: handler,
		brokers: make(map[string]*connect, 16), // 一般不会超过 16 个 broker
		cfg:     cfg,
		random:  random,
	}
	transport := &http.Transport{DialContext: hub.dialContext}
	cli := &http.Client{Transport: transport}
	client := httpclient.NewClient(cli)
	hub.client = client
	hub.proxy = pubrr.Forward(transport, "manager")

	return hub
}

type brkHub struct {
	db      *gorm.DB
	notice  evtrsk.Handler
	handler http.Handler
	mutex   sync.RWMutex
	brokers map[string]*connect
	client  httpclient.Client
	proxy   pubrr.Forwarder
	cfg     conf.Config
	random  *rand.Rand
}

// Auth 鉴权授权管理
func (hub *brkHub) Auth(ident Ident) (Issue, http.Header, error) {
	var issue Issue
	id, secret, inet := ident.ID, ident.Secret, ident.Inet
	if len(inet) == 0 || inet.IsLoopback() {
		return issue, nil, ErrBrokerInet
	}

	// 查询 broker
	var brk model.Broker
	ipv4 := inet.String()
	if err := hub.db.Take(&brk, "id = ? AND inet = ? AND secret = ?", id, ipv4, secret).
		Error; err != nil {
		return issue, nil, ErrBrokerNotFound
	}
	sid := strconv.FormatInt(id, 10)
	if brk.Status || hub.getConn(sid) != nil {
		return issue, nil, ErrBrokerRepeat
	}

	// 随机生成一个 32-64 位的加密密钥
	psz := hub.random.Intn(33) + 32
	passwd := make([]byte, psz)
	_, _ = hub.random.Read(passwd)

	issue.Name, issue.Passwd = brk.Name, passwd
	issue.Listen = Listen{Addr: ":8180"}
	issue.Logger, issue.Database = hub.cfg.Logger, hub.cfg.Database

	return issue, nil, nil
}

func (hub *brkHub) Join(tran net.Conn, ident Ident, issue Issue) error {
	mux := spdy.Server(tran, spdy.WithEncrypt(issue.Passwd))
	//goland:noinspection GoUnhandledErrorResult
	defer mux.Close()

	id := ident.ID
	sid := strconv.FormatInt(id, 10)
	conn := &connect{
		id:    id,
		sid:   sid,
		ident: ident,
		issue: issue,
		mux:   mux,
	}

	if !hub.putConn(conn) { // [上线] 将 connect 存到连接池中
		return ErrBrokerRepeat
	}
	defer hub.delConn(sid) // [下线] 删除连接池中的

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
		PingedAt:   nowAt,
		JoinedAt:   nowAt,
	}
	if err := hub.db.UpdateColumns(tbl).Error; err != nil { // [上线] 修改为在线状态
		return err
	}
	defer func() {
		hub.db.Model(tbl).Update("status", false) // [下线] 修改为离线状态
	}()

	// 通知上线
	inet := ident.Inet.String()
	online := &model.Event{
		MinionID:  id,
		Inet:      inet,
		Msg:       "代理节点上线",
		Typeof:    "broker.online",
		SendAlert: true,
		OccurAt:   now,
		CreatedAt: now,
	}
	_ = hub.notice.Event(online)

	srv := &http.Server{
		Handler: hub.handler,
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
	_ = hub.notice.Event(offline)

	return nil
}

func (hub *brkHub) ResetDB() error {
	return hub.db.Model(&model.Broker{}).
		Where("status = ?", true).
		UpdateColumn("status", false).
		Error
}

func (hub *brkHub) CallB(ctx context.Context, op opurl.URLer, body io.Reader) (*http.Response, error) {
	req := hub.newRequest(ctx, op, body)
	return hub.client.Fetch(req)
}

func (hub *brkHub) JSONB(ctx context.Context, op opurl.URLer, body, reply any) error {
	rd := hub.readJSON(body)
	req := hub.newRequest(ctx, op, rd)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := hub.client.Fetch(req)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(reply)
}

func (hub *brkHub) OnewayB(ctx context.Context, op opurl.URLer, body io.Reader) error {
	res, err := hub.CallB(ctx, op, body)
	if err == nil {
		_ = res.Body.Close()
	}
	return err
}

func (hub *brkHub) CallM(ctx context.Context) {
}

func (hub *brkHub) Forward(op opurl.URLer, w http.ResponseWriter, r *http.Request) {
	hub.proxy.Forward(op, w, r)
}

// getConn 通过 ID 获取连接
func (hub *brkHub) getConn(id string) *connect {
	hub.mutex.RLock()
	conn := hub.brokers[id]
	hub.mutex.RUnlock()

	return conn
}

// putConn 存入一个 broker 连接
// false- 连接已经存在，存入失败
func (hub *brkHub) putConn(conn *connect) bool {
	id := conn.sid
	hub.mutex.Lock()
	_, exist := hub.brokers[id]
	if !exist {
		hub.brokers[id] = conn
	}
	hub.mutex.Unlock()

	return !exist
}

// delConn 删除 broker 连接
func (hub *brkHub) delConn(id string) bool {
	hub.mutex.Lock()
	_, exist := hub.brokers[id]
	if exist {
		delete(hub.brokers, id)
	}
	hub.mutex.Unlock()

	return exist
}

func (hub *brkHub) dialContext(_ context.Context, _, addr string) (net.Conn, error) {
	id, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, net.InvalidAddrError(addr)
	}

	if conn := hub.getConn(id); conn != nil {
		return conn.mux.Dial()
	}

	return nil, ErrBrokerOffline
}

func (*brkHub) newRequest(ctx context.Context, op opurl.URLer, body io.Reader) *http.Request {
	method := op.Method()
	addr := op.URL()
	req := &http.Request{
		Method:     method,
		URL:        addr,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
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
	if le, ok := body.(interface{ Len() int }); ok {
		req.ContentLength = int64(le.Len())
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

func (*brkHub) readJSON(v any) *jsonReader {
	if v == nil {
		return &jsonReader{err: io.EOF}
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(v)
	return &jsonReader{err: err, buf: buf}
}

type jsonReader struct {
	err error
	buf *bytes.Buffer
}

func (jr *jsonReader) Read(p []byte) (int, error) {
	if jr.err != nil {
		return 0, jr.err
	}
	return jr.buf.Read(p)
}

func (jr *jsonReader) Len() int {
	if jr.err != nil || jr.buf == nil {
		return 0
	}
	return jr.buf.Len()
}

func (jr *jsonReader) Close() error {
	return nil
}
