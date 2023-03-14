package restful

import (
	"net/http"

	"github.com/vela-ssoc/vela-manager/restful/brksrv"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

func Brk(sh *ship.Ship, db *gorm.DB) http.Handler {
	group := sh.Group("/api/v1")
	brksrv.Ping(db).Route(group)

	return sh
}
