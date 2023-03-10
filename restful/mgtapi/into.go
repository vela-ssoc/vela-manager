package mgtapi

import (
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/vela-ssoc/backend-common/errno"
	"github.com/vela-ssoc/backend-common/model"
	"github.com/vela-ssoc/backend-common/netutil"
	"github.com/vela-ssoc/backend-common/opurl"
	"github.com/vela-ssoc/vela-manager/blink"
	"github.com/vela-ssoc/vela-manager/restful/facade"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

// Into 节点透传调用接口
func Into(db *gorm.DB, hub blink.Huber, node string) facade.MgtRouter {
	upg := netutil.Upgrade(node)

	return &intoCtrl{
		db:  db,
		hub: hub,
		upg: upg,
	}
}

type intoCtrl struct {
	db  *gorm.DB
	hub blink.Huber
	upg websocket.Upgrader
}

func (ac *intoCtrl) Route(_, auth *ship.RouteGroupBuilder) {
	auth.Route("/brr/:arg/*path").Any(ac.Broker(ac.ForwardB)) // http 穿透
	auth.Route("/arr/:arg/*path").Any(ac.Minion(ac.ForwardA)) // http 穿透
	auth.Route("/bws/:arg/*path").GET(ac.Broker(ac.SocketB))  // socket 穿透
	auth.Route("/aws/:arg/*path").GET(ac.Minion(ac.SocketA))  // socket 穿透
}

func (ac *intoCtrl) Broker(fn func(*ship.Context, string, *model.Broker) error) ship.Handler {
	return func(c *ship.Context) error {
		// arg 参数既可以是节点 ID 也可以是节点 IP，程序需要判断自适应
		arg := c.Param("arg")
		path := c.Param("path")
		bid, _ := strconv.ParseInt(arg, 10, 64)
		var err error
		brk := new(model.Broker)
		tx := ac.db.Select("id", "inet", "status")
		if bid != 0 {
			err = tx.First(brk, "id = ?", bid).Error
		} else {
			err = tx.First(brk, "inet = ?", arg).Error
		}
		if err != nil {
			return errno.NodeNotfound(arg)
		}
		if !brk.Status {
			// return errno.NodeOffline(brk.ID, brk.Inet)
			return errno.NodeOffline(brk.ID, "")
		}

		return fn(c, path, brk)
	}
}

func (ac *intoCtrl) Minion(fn func(*ship.Context, string, *model.Minion) error) ship.Handler {
	return func(c *ship.Context) error {
		// arg 参数既可以是节点 ID 也可以是节点 IP，程序需要判断自适应
		arg := c.Param("arg")
		path := c.Param("path")
		mid, _ := strconv.ParseInt(arg, 10, 64)

		var err error
		mon := new(model.Minion)
		tx := ac.db.Select("id", "inet", "status", "broker_id")
		if mid != 0 {
			err = tx.First(mon, "id = ?", mid).Error
		} else {
			err = tx.First(mon, "inet = ?", arg).Error
		}
		if err != nil {
			return errno.NodeNotfound(arg)
		}
		sta := mon.Status
		if sta == model.MinionInactive {
			return errno.NodeInactive(mon.ID, mon.Inet)
		} else if sta == model.MinionOffline {
			return errno.NodeOffline(mon.ID, mon.Inet)
		} else if sta == model.MinionRemove {
			return errno.NodeRemove(mon.ID, mon.Inet)
		}

		return fn(c, path, mon)
	}
}

func (ac *intoCtrl) ForwardB(c *ship.Context, path string, brk *model.Broker) error {
	w, r := c.ResponseWriter(), c.Request()
	op := opurl.MBrr(brk.ID, c.Method(), path, r.URL.RawQuery)
	ac.hub.Forward(op, w, r)

	return nil
}

func (ac *intoCtrl) ForwardA(c *ship.Context, path string, mon *model.Minion) error {
	w, r := c.ResponseWriter(), c.Request()
	op := opurl.MArr(mon.BrokerID, mon.ID, c.Method(), path, r.URL.RawQuery)
	ac.hub.Forward(op, w, r)

	return nil
}

func (ac *intoCtrl) SocketB(c *ship.Context, path string, brk *model.Broker) error {
	if !c.IsWebSocket() {
		return ship.ErrBadRequest
	}

	w, r := c.ResponseWriter(), c.Request()
	op := opurl.MBws(brk.ID, path, r.URL.RawQuery)
	back, err := ac.hub.Stream(op, nil)
	if err != nil {
		return err
	}

	fore, err := ac.upg.Upgrade(w, r, nil)
	if err != nil {
		_ = back.Close()
		return err
	}

	netutil.Pipe(fore, back)

	return nil
}

func (ac *intoCtrl) SocketA(c *ship.Context, path string, mon *model.Minion) error {
	if !c.IsWebSocket() {
		return ship.ErrBadRequest
	}

	c.Infof("frontend -> manager -> broker -> minion 正在准备建立双向流隧道")

	w, r := c.ResponseWriter(), c.Request()
	op := opurl.MAws(mon.BrokerID, mon.ID, path, r.URL.RawQuery)
	back, err := ac.hub.Stream(op, nil)
	if err != nil {
		c.Warnf("与 minion(%s) 建立 websocket 失败: %v", mon.Inet, err)
		return err
	}

	c.Infof("broker -> minion 段隧道已经建立成功")
	fore, err := ac.upg.Upgrade(w, r, nil)
	if err != nil {
		_ = back.Close()
		c.Warnf("与 frontend -> manager upgrade websocket 失败: %v", mon.Inet, err)
		return err
	}

	c.Infof("frontend -> manager -> broker -> minion  段隧道已经建立成功")
	netutil.Pipe(fore, back)
	c.Infof("frontend -> manager -> broker -> minion 隧道已关闭")

	return nil
}
