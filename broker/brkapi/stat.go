package brkapi

import (
	"github.com/xgfone/ship/v5"
)

func Stat() BindRouter {
	return &statAPI{}
}

type statAPI struct{}

func (st *statAPI) BindRoute(rgb *ship.RouteGroupBuilder) {
	rgb.Route("/stat").POST(st.Stat)
}

func (st *statAPI) Stat(c *ship.Context) error {
	return nil
}
