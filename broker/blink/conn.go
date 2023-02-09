package blink

import (
	"context"

	"github.com/dfcfw/spdy"
)

type Broker interface {
	ID() int64
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
func (c *connect) Ident() Ident { return c.ident }
func (c *connect) Issue() Issue { return c.issue }

func Ctx(ctx context.Context) Broker {
	if ctx != nil {
		brk := ctx.Value(brokerCtxKey).(Broker)
		return brk
	}

	return nil
}

type contextKey struct{ name string }

var brokerCtxKey = &contextKey{name: "broker-context"}
