package service

import (
	"context"

	"github.com/vela-ssoc/vela-manager/bridge/linkhub"
)

func NewBrokerCommand(hub linkhub.Huber) *BrokerCommand {
	return &BrokerCommand{
		hub: hub,
	}
}

type BrokerCommand struct {
	hub linkhub.Huber
}

// Exit 让 broker 退出执行。
// 重启 broker 时可以用到，一般 broker 退出后，systemctl 会拉起 broker 程序，
// 从而实现重启的效果。
func (bc *BrokerCommand) Exit(ctx context.Context, id int64) error {
	_ = bc.hub.Oneway(ctx, id, "/api/v1/system/exit", nil)
	return nil
}
