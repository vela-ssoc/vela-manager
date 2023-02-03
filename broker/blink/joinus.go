package blink

import (
	"context"
	"database/sql"
	"errors"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/dfcfw/spdy"
	"github.com/vela-ssoc/manager/inward/evtrsk"
	"github.com/vela-ssoc/manager/model"
	"gorm.io/gorm"
)

var (
	ErrBrokerNotFound = errors.New("broker 节点不存在")
	ErrBrokerRepeat   = errors.New("broker 节点重复连接")
	ErrBrokerInet     = errors.New("broker IP 不合法")
)

type Huber interface {
	Joiner
	Reset()
}

// Hub broker 节点的连接中心
func Hub(db *gorm.DB, notice evtrsk.Handler, handler http.Handler) Huber {
	return &brkHub{
		db:      db,
		notice:  notice,
		handler: handler,
		brokers: make(map[int64]*connect, 16), // 一般不会超过 16 个 broker
	}
}

type brkHub struct {
	db      *gorm.DB
	notice  evtrsk.Handler
	handler http.Handler
	mutex   sync.RWMutex
	brokers map[int64]*connect
}

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

	return Grant{}, nil, nil
}

func (bh *brkHub) Join(tran net.Conn, ident Ident, ret any) error {
	grant := ret.(Grant)
	mux := spdy.Server(tran)

	id := ident.ID
	conn := &connect{
		ident:  ident,
		grant:  grant,
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
