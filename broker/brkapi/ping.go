package brkapi

import (
	"time"

	"github.com/vela-ssoc/manager/broker/blink"
	"github.com/vela-ssoc/manager/model"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

func Ping(db *gorm.DB) BindRouter {
	return &pingAPI{db: db}
}

type pingAPI struct {
	db *gorm.DB
}

func (pg *pingAPI) BindRoute(rgb *ship.RouteGroupBuilder) {
	rgb.Route("/ping").GET(pg.Ping)
}

func (pg *pingAPI) Ping(c *ship.Context) error {
	brk := blink.Ctx(c.Request().Context())
	tbl := &model.Broker{ID: brk.ID()}
	pg.db.Model(tbl).Update("ping_at", time.Now())
	return nil
}
