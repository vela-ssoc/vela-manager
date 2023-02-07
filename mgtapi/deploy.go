package mgtapi

import (
	"github.com/vela-ssoc/manager/infra/grid"
	"github.com/vela-ssoc/manager/model"
	"github.com/xgfone/ship/v5"
	"gorm.io/gorm"
)

func Deploy(db *gorm.DB, gfs grid.FS) RouteBinder {
	return &deployCtrl{db: db, gfs: gfs}
}

type deployCtrl struct {
	db  *gorm.DB // 数据库连接
	gfs grid.FS  // 文件管理系统
}

func (dep *deployCtrl) BindRoute(anon, _ *ship.RouteGroupBuilder) {
	anon.Route("/deploy").GET(dep.Broker)
}

func (dep *deployCtrl) Broker(c *ship.Context) error {
	type request struct {
		Arch   string       `form:"arch"   validate:"oneof=386 amd64 arm arm64"`
		Goos   string       `form:"goos"   validate:"oneof=linux windows darwin"`
		Semver model.Semver `form:"semver"`
	}
	var req request
	if err := c.BindQuery(&req); err != nil {
		return err
	}

	return nil
}
