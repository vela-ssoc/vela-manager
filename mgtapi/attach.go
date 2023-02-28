package mgtapi

import (
	"net"

	"github.com/gorilla/websocket"
	"github.com/vela-ssoc/backend-common/model"
	"github.com/vela-ssoc/backend-common/netutil"
	"github.com/vela-ssoc/backend-common/opurl"
	"github.com/vela-ssoc/vela-manager/blink"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

func Attach(db *gorm.DB, hub blink.Huber, node string) RouteBinder {
	upg := netutil.Upgrade(node)

	return &attachCtrl{
		db:  db,
		hub: hub,
		upg: upg,
	}
}

type attachCtrl struct {
	db  *gorm.DB
	hub blink.Huber
	upg websocket.Upgrader
}

func (ac *attachCtrl) BindRoute(_, auth *ship.RouteGroupBuilder) {
	auth.Route("/brr/:arg/*path").Any(ac.Broker(ac.ForwardB)) // http 穿透
	auth.Route("/mrr/:arg/*path").Any(ac.Minion(ac.ForwardM)) // http 穿透
	auth.Route("/bws/:arg/*path").GET(ac.Broker(ac.SocketB))  // socket 穿透
	auth.Route("/mws/:arg/*path").GET(ac.Minion(ac.SocketM))  // socket 穿透
}

func (ac *attachCtrl) Broker(fn func(*ship.Context, string, *model.Broker) error) ship.Handler {
	return func(c *ship.Context) error {
		// param 参数既可以是节点 ID 也可以是节点 IP，程序需要判断自适应
		param := c.Param("arg")
		path := c.Param("path")
		ipv4 := net.ParseIP(param).To4()
		var err error
		brk := new(model.Broker)
		tx := ac.db.Select("id", "inet", "status")
		if ipv4 != nil {
			err = tx.First(brk, "inet = ?", ipv4.String()).Error
		} else {
			err = tx.First(brk, "id = ?", param).Error
		}
		if err != nil {
			return err
		}

		return fn(c, path, brk)
	}
}

func (ac *attachCtrl) Minion(fn func(*ship.Context, string, *model.Minion) error) ship.Handler {
	return func(c *ship.Context) error {
		// arg 参数既可以是节点 ID 也可以是节点 IP，程序需要判断自适应
		arg := c.Param("arg")
		path := c.Param("path")
		ipv4 := net.ParseIP(arg).To4()

		var err error
		mon := new(model.Minion)
		tx := ac.db.Select("id", "inet", "status", "broker_id")
		if ipv4 != nil {
			err = tx.First(mon, "inet = ?", ipv4.String()).Error
		} else {
			err = tx.First(mon, "id = ?", arg).Error
		}
		if err != nil {
			return err
		}

		return fn(c, path, mon)
	}
}

func (ac *attachCtrl) ForwardB(c *ship.Context, path string, brk *model.Broker) error {
	w, r := c.ResponseWriter(), c.Request()
	op := opurl.MBrr(brk.ID, c.Method(), path, r.URL.RawQuery)
	ac.hub.Forward(op, w, r)

	return nil
}

func (ac *attachCtrl) ForwardM(c *ship.Context, path string, mon *model.Minion) error {
	w, r := c.ResponseWriter(), c.Request()
	op := opurl.MMrr(mon.BrokerID, mon.ID, c.Method(), path, r.URL.RawQuery)
	ac.hub.Forward(op, w, r)

	return nil
}

func (ac *attachCtrl) SocketB(c *ship.Context, path string, brk *model.Broker) error {
	if !c.IsWebSocket() {
		return nil
	}

	w, r := c.ResponseWriter(), c.Request()
	op := opurl.MBws(brk.ID, path, r.URL.RawQuery)
	back, err := ac.hub.Stream(op, nil)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer back.Close()

	fore, err := ac.upg.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer fore.Close()

	netutil.Pipe(fore, back)

	return nil
}

func (ac *attachCtrl) SocketM(c *ship.Context, path string, mon *model.Minion) error {
	if !c.IsWebSocket() {
		return ship.ErrTooManyRequests
	}

	w, r := c.ResponseWriter(), c.Request()
	op := opurl.MMws(mon.BrokerID, mon.ID, path, r.URL.RawQuery)
	back, err := ac.hub.Stream(op, nil)
	if err != nil {
		c.Warnf("与 minion(%s) 建立 websocket 失败: %v", mon.Inet, err)
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer back.Close()
	c.Warnf("与 minion(%s) websocket 隧道已打通", mon.Inet)

	fore, err := ac.upg.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer fore.Close()

	netutil.Pipe(fore, back)

	return nil
}

// FM File Manager
func (ac *attachCtrl) FM(c *ship.Context) error {
	return nil
}
