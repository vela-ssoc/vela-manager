package mgtapi

import (
	"net/http"

	"github.com/vela-ssoc/manager/broker/blink"
	"github.com/xgfone/ship/v5"
)

func Attach(hub blink.Huber) RouteBinder {
	return &attachCtrl{hub: hub}
}

type attachCtrl struct {
	hub blink.Huber
}

func (ac *attachCtrl) BindRoute(anon, _ *ship.RouteGroupBuilder) {
	anon.Route("/broker/attach/environ").GET(ac.Environ)
}

func (ac *attachCtrl) Environ(c *ship.Context) error {
	var req struct {
		ID int64 `query:"id"`
	}
	if err := c.BindQuery(&req); err != nil {
		return err
	}

	res, err := ac.hub.Fetch(nil, req.ID, blink.BrkEnv, nil)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()
	ct := res.Header.Get(ship.HeaderContentType)

	return c.Stream(http.StatusOK, ct, res.Body)
}
