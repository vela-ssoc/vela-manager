package mgtapi

import "github.com/xgfone/ship/v5"

// RouteBinder 路由绑定接口
type RouteBinder interface {
	// BindRoute anon 为无需认证即可访问的接口
	// auth 为必须登录认证才可访问的接口
	BindRoute(anon, auth *ship.RouteGroupBuilder)
}
