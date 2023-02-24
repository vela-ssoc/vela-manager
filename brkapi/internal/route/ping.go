package route

import (
	"time"

	"github.com/vela-ssoc/backend-common/model"
	"github.com/vela-ssoc/vela-manager/blink"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

func Ping(db *gorm.DB) RegRouter {
	return &pingCtrl{db: db}
}

type pingCtrl struct {
	db *gorm.DB
}

func (pc *pingCtrl) RegRoute(rgb *ship.RouteGroupBuilder) {
	rgb.Route("/ping").GET(pc.Ping)
}

func (pc *pingCtrl) Ping(c *ship.Context) error {
	infer := blink.Ctx(c.Request().Context())
	bid := infer.ID()

	return pc.db.Model(&model.Broker{ID: bid}).
		Where("status = ?", true).
		UpdateColumn("pinged_at", time.Now()).
		Error
}
