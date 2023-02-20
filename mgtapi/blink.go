package mgtapi

import (
	"net/http"

	"github.com/xgfone/ship/v5"
)

// Blink broker 节点接入层
func Blink(join http.Handler) RouteBinder {
	return &blinkCtrl{join: join}
}

type blinkCtrl struct {
	join http.Handler
}

func (bc *blinkCtrl) BindRoute(anon, _ *ship.RouteGroupBuilder) {
	anon.Route("/broker").CONNECT(bc.Join)
}

func (bc *blinkCtrl) Join(c *ship.Context) error {
	bc.join.ServeHTTP(c.ResponseWriter(), c.Request())
	return nil
}
