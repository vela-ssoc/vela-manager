package middle

import (
	"net/http"
	"strings"

	"github.com/xgfone/ship/v5"
)

// Auth 登录认证中间件
func Auth() ship.Middleware {
	ath := &auth{
		key: strings.ToLower(ship.HeaderAuthorization),
	}
	return ath.middleware
}

type auth struct {
	key string // Token 的 key
}

// middleware 认证中间件
func (ath *auth) middleware(handler ship.Handler) ship.Handler {
	return func(c *ship.Context) error {
		// 从 Header 获取 Token，当 Header 中为找到 Token 后再从查询参数中读取，
		// 只允许 GET 方法从 header 中获取，为了照顾下载请求和 websocket
		token := c.GetReqHeader(ath.key)
		if token == "" && c.Method() == http.MethodGet {
			token = c.Query(ath.key)
		}
		// TODO: 拦截认证无效的请求
		if val, err := c.GetSession(token); err == nil {
			c.Any = val
		}

		return handler(c)
	}
}
