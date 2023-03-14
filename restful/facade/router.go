package facade

import "github.com/xgfone/ship/v5"

// MgtRouter manager 路由注册绑定
type MgtRouter interface {
	// Route 绑定路由
	// anon 为无需认证即可访问的接口
	// auth 为必须登录认证才可访问的接口
	Route(anon, auth *ship.RouteGroupBuilder)
}

// BrkRouter 为 broker 提供的接口
type BrkRouter interface {
	// Route 绑定路由
	Route(*ship.RouteGroupBuilder)
}
