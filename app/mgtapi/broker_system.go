package mgtapi

import (
	"github.com/vela-ssoc/vela-common-mb/param/request"
	"github.com/vela-ssoc/vela-manager/app/route"
	"github.com/vela-ssoc/vela-manager/app/service"
	"github.com/xgfone/ship/v5"
)

func NewBrokerSystem(svc *service.BrokerSystem) *BrokerSystem {
	return &BrokerSystem{
		svc: svc,
	}
}

type BrokerSystem struct {
	svc *service.BrokerSystem
}

func (bc *BrokerSystem) Route(_, bearer, _ *ship.RouteGroupBuilder) {
	bearer.Route("/broker/command/exit").Data(route.Ignore()).GET(bc.exit)
}

// Exit 让 broker 退出执行。
// 重启 broker 时可以用到，一般 broker 退出后，systemctl 会拉起 broker 程序，
// 从而实现重启的效果。
func (bc *BrokerSystem) exit(c *ship.Context) error {
	req := new(request.Int64ID)
	if err := c.BindQuery(req); err != nil {
		return err
	}
	ctx := c.Request().Context()

	return bc.svc.Exit(ctx, req.ID)
}
