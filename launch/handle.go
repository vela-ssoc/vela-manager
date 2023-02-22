package launch

import (
	"net/http"

	"github.com/xgfone/ship/v5"
)

func handleError(c *ship.Context, err error) {
	c.WriteHeader(http.StatusBadRequest)
}

type result struct {
	Code  int    `json:"code"`
	Cause string `json:"cause"`
}
