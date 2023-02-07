package hanerr

import (
	"net/http"
	"time"

	"github.com/xgfone/ship/v5"
)

func Handle(c *ship.Context, err error) {
	ret := &resultError{Code: http.StatusBadRequest}
	switch e := err.(type) {
	case ship.HTTPServerError:
		ret.Code, ret.Cause = e.Code, e.Error()
	case *time.ParseError:
		ret.Cause = "时间格式化错误（正确格式：" + e.Layout + "）"
	default:
		ret.Cause = e.Error()
	}
	code := ret.Code
	if code < http.StatusBadRequest {
		code = http.StatusBadRequest
	}
	_ = c.JSON(code, ret)
}

type resultError struct {
	Code  int    `json:"code"`  // 错误码
	Cause string `json:"cause"` // 错误原因
}
