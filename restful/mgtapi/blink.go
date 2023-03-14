package mgtapi

import (
	"net/http"

	"github.com/vela-ssoc/vela-manager/restful/facade"
	"github.com/xgfone/ship/v5"
)

func Blink(gw http.Handler) facade.MgtRouter {
	return &blinkCtrl{gw: gw}
}

type blinkCtrl struct {
	gw http.Handler
}

func (bl *blinkCtrl) Route(anon, _ *ship.RouteGroupBuilder) {
	anon.Route("/broker").CONNECT(bl.Join)
}

func (bl *blinkCtrl) Join(c *ship.Context) error {
	w, r := c.ResponseWriter(), c.Request()
	bl.gw.ServeHTTP(w, r)
	return nil
}
