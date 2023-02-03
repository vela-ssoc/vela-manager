package blink

import (
	"context"

	"github.com/dfcfw/spdy"
)

type Broker interface {
	ID() int64
	Ident() Ident
	Grant() Grant
}

type connect struct {
	ident  Ident
	grant  Grant
	mux    spdy.Muxer
	waiter *brkHub
}

func (c *connect) ID() int64    { return c.ident.ID }
func (c *connect) Ident() Ident { return c.ident }
func (c *connect) Grant() Grant { return c.grant }

func Ctx(ctx context.Context) Broker {
	if ctx != nil {
		brk := ctx.Value(brokerCtxKey).(Broker)
		return brk
	}

	return nil
}

type contextKey struct{ name string }

var brokerCtxKey = &contextKey{name: "broker-context"}
