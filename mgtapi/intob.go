package mgtapi

import (
	"net"

	"github.com/vela-ssoc/backend-common/model"
	"github.com/vela-ssoc/backend-common/opurl"
	"github.com/vela-ssoc/vela-manager/blink"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

func Intob(db *gorm.DB, hub blink.Huber) RouteBinder {
	return &intob{
		db:  db,
		hub: hub,
	}
}

type intob struct {
	db  *gorm.DB
	hub blink.Huber
}

func (inb *intob) BindRoute(anon, auth *ship.RouteGroupBuilder) {
	anon.Route("/intob/:param/ping").GET(inb.Warp(inb.Ping))
	anon.Route("/intob/:param/syscmd").GET(inb.Warp(inb.Syscmd))
	anon.Route("/intob/:param/fs").GET(inb.Warp(inb.FS))
}

func (inb *intob) Warp(fn func(*ship.Context, *model.Broker) error) ship.Handler {
	return func(c *ship.Context) error {
		// param 参数既可以是节点 ID 也可以是节点 IP，程序需要判断自适应
		param := c.Param("param")
		ipv4 := net.ParseIP(param).To4()

		var err error
		brk := new(model.Broker)
		tx := inb.db.Select("id", "inet", "status")
		if ipv4 != nil {
			err = tx.First(brk, "inet = ?", ipv4.String()).Error
		} else {
			err = tx.First(brk, "id = ?", param).Error
		}
		if err != nil {
			return err
		}

		return fn(c, brk)
	}
}

func (inb *intob) Ping(c *ship.Context, brk *model.Broker) error {
	c.Infof("ping broker: %s", brk.Inet)
	op := opurl.OpPing.IntID(brk.ID)
	return inb.hub.OnewayB(nil, op, nil)
}

// Syscmd 执行系统命令
func (inb *intob) Syscmd(c *ship.Context, brk *model.Broker) error {
	query := c.Request().URL.RawQuery
	op := opurl.OpIntobSyscmd.IntID(brk.ID).WithQuery(query)

	inb.hub.ForwardB(op, c.ResponseWriter(), c.Request())

	return nil
}

// FS 文件系统
func (inb *intob) FS(c *ship.Context, brk *model.Broker) error {
	query := c.Request().URL.RawQuery
	op := opurl.BrkFS.IntID(brk.ID).WithQuery(query)

	inb.hub.ForwardB(op, c.ResponseWriter(), c.Request())

	return nil
}
