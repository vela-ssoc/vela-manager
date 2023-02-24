package blink

import (
	"context"
	"net"
)

// Infer broker 节点信息接口
type Infer interface {
	// ID broker 节点 ID
	ID() int64

	// Name broker 节点名字
	Name() string

	// Inet broker 节点 IP 地址
	Inet() net.IP

	// Ident broker 节点上线认证时自身收集的认证信息
	Ident() Ident

	// Issue broker 上线认证成功后下发的信息
	Issue() Issue
}

func Ctx(ctx context.Context) Infer {
	if ctx != nil {
		brk := ctx.Value(brokerCtxKey).(Infer)
		return brk
	}

	return nil
}

type contextKey struct{ name string }

var brokerCtxKey = &contextKey{name: "broker-context"}
