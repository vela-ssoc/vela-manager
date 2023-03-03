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

	"github.com/gorilla/websocket"
	"github.com/vela-ssoc/backend-common/logback"
	"github.com/vela-ssoc/backend-common/model"
	"github.com/vela-ssoc/backend-common/netutil"
	"github.com/vela-ssoc/backend-common/opurl"
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

// Huber broker 接入中心
type Huber interface {
	// Joiner broker 上线认证加入接口
	Joiner

	// ResetDB 将数据库中所有在线的 broker 修改为离线，该接口不会清除 hub 中的连接池，
	// 故：此方法只适用于程序刚启动时和程序关闭时节点状态归位。
	ResetDB() error

	// Fetch 请求调用 broker 节点
	Fetch(context.Context, opurl.URLer, io.Reader) (*http.Response, error)

	// JSON 请求调用 broker 节点，请求和响应的数据会进行 json 序列化和反序列化。
	JSON(context.Context, opurl.URLer, any, any) error

	// Oneway 请求调用 broker 节点，但是不会解析 broker 的响应结果。
	Oneway(context.Context, opurl.URLer, io.Reader) error

	// Forward 代理模式请求响应。
	Forward(opurl.URLer, http.ResponseWriter, *http.Request)

	// Stream 建立 websocket.Conn 双向流
	Stream(opurl.URLer, http.Header) (*websocket.Conn, error)
}

// Hub broker 节点的连接中心
func Hub(db *gorm.DB, notice evtrsk.Handler, handler http.Handler, cfg conf.Config, slog logback.Logger, node string) Huber {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	hub := &brkHub{
		db:      db,
		notice:  notice,
		slog:    slog,
		handler: handler,
		brokers: make(map[string]*connect, 16), // 一般不会超过 16 个 broker
		cfg:     cfg,
		random:  random,
	}
	transport := &http.Transport{DialContext: hub.dialContext}
	hub.client = opurl.NewClient(transport, slog)
	hub.proxy = netutil.Forward(transport, node)
	hub.stream = netutil.Stream(hub.dialContext)

	return hub
}

type brkHub struct {
	db      *gorm.DB
	slog    logback.Logger
	notice  evtrsk.Handler
	handler http.Handler
	mutex   sync.RWMutex
	brokers map[string]*connect
	client  opurl.Client
	proxy   netutil.Forwarder
	stream  netutil.Streamer
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

// ResetDB 将数据库中在线的 broker 节点修改为离线
func (hub *brkHub) ResetDB() error {
	return hub.db.Model(&model.Broker{}).
		Where("status = ?", true).
		UpdateColumn("status", false).
		Error
}

// Fetch 向 broker 节点发送请求
func (hub *brkHub) Fetch(ctx context.Context, op opurl.URLer, body io.Reader) (*http.Response, error) {
	return hub.client.Fetch(ctx, op, nil, body)
}

// JSON 向 broker 发送 JSON 格式的请求
func (hub *brkHub) JSON(ctx context.Context, op opurl.URLer, body, reply any) error {
	return hub.client.JSON(ctx, op, nil, body, reply)
}

// Oneway 向 broker 发送请求不关心 broker 响应内容
func (hub *brkHub) Oneway(ctx context.Context, op opurl.URLer, body io.Reader) error {
	res, err := hub.client.Fetch(ctx, op, nil, body)
	if err == nil {
		_ = res.Body.Close()
	}
	return err
}

// Forward 代理模式发送请求
func (hub *brkHub) Forward(op opurl.URLer, w http.ResponseWriter, r *http.Request) {
	hub.proxy.Forward(op, w, r)
}

// Stream 通过 websocket 方式建立双向流
func (hub *brkHub) Stream(op opurl.URLer, header http.Header) (*websocket.Conn, error) {
	addr := op.String()
	conn, _, err := hub.stream.Stream(addr, header)
	if err == nil {
		hub.slog.Infof("建立 stream (%s) 通道成功", addr)
	} else {
		hub.slog.Warnf("建立 stream (%s) 通道失败：%s", addr, err)
	}
	return conn, err
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

func (*brkHub) newRequest(ctx context.Context, op opurl.URLer, header http.Header, body io.Reader) *http.Request {
	method, dst := op.Method(), op.URL()
	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = io.NopCloser(body)
	}
	req := &http.Request{
		Method:     method,
		URL:        dst,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     header,
		Body:       rc,
	}
	if req.Header == nil {
		req.Header = make(http.Header)
	}
	if body != nil {
		switch v := body.(type) {
		case *bytes.Buffer:
			req.ContentLength = int64(v.Len())
			buf := v.Bytes()
			req.GetBody = func() (io.ReadCloser, error) {
				r := bytes.NewReader(buf)
				return io.NopCloser(r), nil
			}
		case *bytes.Reader:
			req.ContentLength = int64(v.Len())
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return io.NopCloser(&r), nil
			}
		case *strings.Reader:
			req.ContentLength = int64(v.Len())
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return io.NopCloser(&r), nil
			}
		case *jsonBody:
			req.ContentLength = int64(v.Len())
			req.GetBody = func() (io.ReadCloser, error) {
				return v, nil
			}
		default:
			// This is where we'd set it to -1 (at least
			// if body != NoBody) to mean unknown, but
			// that broke people during the Go 1.8 testing
			// period. People depend on it being 0 I
			// guess. Maybe retry later. See Issue 18117.
		}
		// For client requests, Request.ContentLength of 0
		// means either actually 0, or unknown. The only way
		// to explicitly say that the ContentLength is zero is
		// to set the Body to nil. But turns out too much code
		// depends on NewRequest returning a non-nil Body,
		// so we use a well-known ReadCloser variable instead
		// and have the http package also treat that sentinel
		// variable to mean explicitly zero.
		if req.GetBody != nil && req.ContentLength == 0 {
			req.Body = http.NoBody
			req.GetBody = func() (io.ReadCloser, error) { return http.NoBody, nil }
		}
	}
	if ctx == nil {
		ctx = context.Background()
	}
	return req.WithContext(ctx)
}

func (*brkHub) jsonBody(v any) *jsonBody {
	if v == nil {
		return nil
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(v)
	return &jsonBody{err: err, buf: buf}
}

type jsonBody struct {
	err error
	buf *bytes.Buffer
}

func (jb *jsonBody) Read(p []byte) (int, error) {
	if jb.err != nil {
		return 0, jb.err
	}
	if jb.buf == nil {
		return 0, io.EOF
	}
	return jb.buf.Read(p)
}

func (jb *jsonBody) Close() error { return nil }

func (jb *jsonBody) Len() int {
	if jb.err != nil || jb.buf == nil {
		return 0
	}
	return jb.buf.Len()
}
