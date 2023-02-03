package brkapi

import "github.com/xgfone/ship/v5"

type BindRouter interface {
	BindRoute(*ship.RouteGroupBuilder)
}
