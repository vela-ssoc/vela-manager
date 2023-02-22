package broker

import (
	"net/http"

	"github.com/vela-ssoc/backend-common/logback"
	"github.com/vela-ssoc/backend-common/validate"
	"github.com/vela-ssoc/vela-manager/broker/brkapi"
	"github.com/vela-ssoc/vela-manager/broker/hanerr"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

func New(db *gorm.DB, valid validate.Validator, logger logback.Logger) http.Handler {
	sh := ship.Default()
	sh.HandleError = hanerr.Func
	sh.Logger = logger
	sh.Validator = valid

	rout := sh.Group("/api")
	brkapi.Ping(db).BindRoute(rout)
	brkapi.Stat().BindRoute(rout)

	return sh
}
