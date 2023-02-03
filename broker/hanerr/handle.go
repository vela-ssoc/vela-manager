package hanerr

import (
	"net/http"

	"github.com/xgfone/ship/v5"
)

type result struct {
	Message string `json:"message"`
}

func Func(c *ship.Context, err error) {
	var msg string
	switch ex := err.(type) {
	default:
		msg = ex.Error()
	}
	res := &result{Message: msg}

	_ = c.JSON(http.StatusBadRequest, res)
}
