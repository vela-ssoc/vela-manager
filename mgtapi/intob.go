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
	anon.Route("/intob/:param/*path").Any(inb.Into)
}

func (inb *intob) Into(c *ship.Context) error {
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

	w, r := c.ResponseWriter(), c.Request()
	path := c.Param("path")
	op := opurl.Intob(brk.ID, c.Method(), path, r.URL.RawQuery)

	inb.hub.Forward(op, w, r)

	return nil
}
