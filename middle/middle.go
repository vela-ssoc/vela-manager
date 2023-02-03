package middle

import "github.com/xgfone/ship/v5"

func Auth1(handler ship.Handler) ship.Handler {
	return func(c *ship.Context) error {
		return handler(c)
	}
}
