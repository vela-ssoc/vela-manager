package blink

import (
	"net"

	"github.com/vela-ssoc/backend-common/spdy"
)

// connect broker 连接结构体
type connect struct {
	id    int64      // broker id
	sid   string     // broker string id
	ident Ident      // 自身认证信息
	issue Issue      // 下发的信息
	mux   spdy.Muxer // 多路复用连接
}

func (c *connect) ID() int64    { return c.ident.ID }   // ID 返回 broker 节点的 ID
func (c *connect) Name() string { return c.issue.Name } // Name broker 节点的名字
func (c *connect) Inet() net.IP { return c.ident.Inet } // Inet broker 的出口 IP
func (c *connect) Ident() Ident { return c.ident }      // Ident broker 节点的身份认证信息
func (c *connect) Issue() Issue { return c.issue }      // Issue 给 broker 下发的授权信息
