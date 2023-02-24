package mgtapi

import (
	"net"

	"github.com/vela-ssoc/backend-common/model"
	"github.com/vela-ssoc/backend-common/opurl"
	"github.com/vela-ssoc/vela-manager/blink"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

func Intom(db *gorm.DB, hub blink.Huber) RouteBinder {
	return &intom{db: db, hub: hub}
}

type intom struct {
	db  *gorm.DB
	hub blink.Huber
}

func (inm *intom) BindRoute(anon, auth *ship.RouteGroupBuilder) {
	anon.Route("/intom/:param/*path").Any(inm.Into)
}

func (inm *intom) Into(c *ship.Context) error {
	// param 参数既可以是节点 ID 也可以是节点 IP，程序需要判断自适应
	param := c.Param("param")
	ipv4 := net.ParseIP(param).To4()

	var err error
	mon := new(model.Minion)
	tx := inm.db.Select("id", "inet", "status", "broker_id")
	if ipv4 != nil {
		err = tx.First(mon, "inet = ?", ipv4.String()).Error
	} else {
		err = tx.First(mon, "id = ?", param).Error
	}
	if err != nil {
		return err
	}

	w, r := c.ResponseWriter(), c.Request()
	query := r.URL.RawQuery
	path := c.Param("path")

	op := opurl.MIntom(mon.BrokerID, mon.ID, c.Method(), path, query)

	inm.hub.Forward(op, w, r)

	return nil
}
