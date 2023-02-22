package brkapi

import (
	"net/http"

	"github.com/vela-ssoc/backend-common/logback"
	"github.com/vela-ssoc/vela-manager/brkapi/internal/route"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

func Handler(db *gorm.DB, valid ship.Validator, slog logback.Logger) http.Handler {
	sh := ship.Default()
	sh.HandleError = handleError
	sh.Validator = valid
	sh.Logger = slog
	sh.NotFound = notfound

	group := sh.Group("/api")
	route.Ping(db).RegRoute(group)

	return sh
}

func handleError(c *ship.Context, err error) {
	code := http.StatusBadRequest
	cause := err.Error()

	res := &result{Code: code, Cause: cause}

	_ = c.JSON(code, res)
}

func notfound(c *ship.Context) error {
	res := &result{
		Code:  http.StatusNotFound,
		Cause: "not found",
	}
	return c.JSON(http.StatusNotFound, res)
}

// result 错误返回结果
type result struct {
	Code  int    `json:"code"`  // 错误码
	Cause string `json:"cause"` // 错误原因
}
