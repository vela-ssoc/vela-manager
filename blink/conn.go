package blink

import (
	"context"
	"net"

	"github.com/vela-ssoc/backend-common/spdy"
)

type Infer interface {
	ID() int64
	Name() string
	Inet() net.IP
	Ident() Ident
	Issue() Issue
}

type connect struct {
	ident  Ident
	issue  Issue
	mux    spdy.Muxer
	waiter *brkHub
}

func (c *connect) ID() int64    { return c.ident.ID }
func (c *connect) Name() string { return c.issue.Name }
func (c *connect) Inet() net.IP { return c.ident.Inet }
func (c *connect) Ident() Ident { return c.ident }
func (c *connect) Issue() Issue { return c.issue }

func Ctx(ctx context.Context) Infer {
	if ctx != nil {
		brk := ctx.Value(brokerCtxKey).(Infer)
		return brk
	}

	return nil
}

type contextKey struct{ name string }

var brokerCtxKey = &contextKey{name: "broker-context"}
