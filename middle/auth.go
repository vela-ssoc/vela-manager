package middle

import (
	"net/http"
	"strings"

	"github.com/xgfone/ship/v5"
)

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
		// 优先从 Header 中读取 Token，找不到再从查询参数中读取，
		// GET 方法从 query 参数中读取 Token 是为了考虑文件下载情况。
		token := c.GetReqHeader(ath.key)
		if token == "" && c.Method() == http.MethodGet {
			token = c.Query(ath.key)
		}
		val, err := c.GetSession(token)
		if err != nil || val == nil {
			return err
		}
		c.Any = val

		return handler(c)
	}
}
