package mgtapi

import (
	"net/http"

	"github.com/xgfone/ship/v5"
)

func Ping() RouteBinder {
	return new(pingCtrl)
}

type pingCtrl struct{}

func (pc *pingCtrl) BindRoute(anon, _ *ship.RouteGroupBuilder) {
	anon.Route("/ping").GET(pc.Ping)
}

func (pc *pingCtrl) Ping(c *ship.Context) error {
	return c.Text(http.StatusOK, "PONG")
}
