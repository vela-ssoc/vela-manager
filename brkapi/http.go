package brkapi

import (
	"net/http"

	"github.com/vela-ssoc/backend-common/logback"
	"github.com/vela-ssoc/backend-common/netutil"
	"github.com/vela-ssoc/vela-manager/brkapi/internal/route"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

func Handler(db *gorm.DB, valid ship.Validator, slog logback.Logger) http.Handler {
	node := "manager"
	sh := ship.Default()
	sh.HandleError = netutil.ErrorFunc(node)
	sh.NotFound = netutil.Notfound(node)
	sh.Validator = valid
	sh.Logger = slog

	group := sh.Group("/api/v1")
	route.Ping(db).RegRoute(group)

	return sh
}
