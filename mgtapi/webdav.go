package mgtapi

import (
	"net/http"
	"strconv"

	"github.com/vela-ssoc/manager/broker/blink"
	"github.com/vela-ssoc/manager/model"
	"github.com/xgfone/ship/v5"
	"golang.org/x/net/webdav"
)

func WebDAV(dir string, hub blink.Huber) RouteBinder {
	if dir == "" {
		dir = "/"
	}
	h := &webdav.Handler{
		FileSystem: webdav.Dir(dir),
		LockSystem: webdav.NewMemLS(),
	}

	return &webdavCtrl{h: h, hub: hub}
}

type webdavCtrl struct {
	h   http.Handler
	hub blink.Huber
}

func (wdc *webdavCtrl) BindRoute(anon, _ *ship.RouteGroupBuilder) {
	methods := []string{
		http.MethodOptions, http.MethodGet, http.MethodHead, http.MethodPost, http.MethodDelete,
		http.MethodPut, "MKCOL", "COPY", "MOVE", "LOCK", "UNLOCK", "PROPFIND", "PROPPATCH",
	}

	dav := anon.Group("/webdav", wdc.BasicAuth) // webdav 仅支持 BasicAuth 认证
	dav.Route("/manager").Method(wdc.Manager, methods...)
	dav.Route("/manager/*path").Method(wdc.Manager, methods...)
	dav.Route("/broker/:bid").Method(wdc.Broker, methods...)
	dav.Route("/broker/:bid/*path").Method(wdc.Broker, methods...)
	dav.Route("/broker/:mid").Method(wdc.Broker, methods...)
	dav.Route("/broker/:mid/*path").Method(wdc.Broker, methods...)
}

// BasicAuth 中间件，webdav 只支持 BasicAuth 认证授权，
// 此处改造成用户名是 username，密码是当前生效的 token
func (wdc *webdavCtrl) BasicAuth(handler ship.Handler) ship.Handler {
	return func(c *ship.Context) error {
		req := c.Request()
		uname, passwd, _ := req.BasicAuth()
		val, err := c.GetSession(passwd)
		info, ok := val.(*model.Userinfo)
		if err != nil || !ok || info.Username != uname {
			c.SetRespHeader(ship.HeaderWWWAuthenticate, `Basic realm="Restricted"`)
			c.WriteHeader(http.StatusUnauthorized)
			return nil
		}

		return handler(c)
	}
}

// Manager 端 webdav 服务
func (wdc *webdavCtrl) Manager(c *ship.Context) error {
	req := c.Request()
	uname, passwd, _ := req.BasicAuth()
	val, err := c.GetSession(passwd)
	info, ok := val.(*model.Userinfo)
	if err != nil || !ok || info.Username != uname {
		c.SetRespHeader(ship.HeaderWWWAuthenticate, `Basic realm="Restricted"`)
		c.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	path := c.Param("path")
	req.URL.Path = "/" + path

	wdc.h.ServeHTTP(c.ResponseWriter(), req)

	return nil
}

func (wdc *webdavCtrl) Broker(c *ship.Context) error {
	bid := c.Param("bid")
	brkID, err := strconv.ParseInt(bid, 10, 64)
	if err != nil {
		return err
	}

	path := c.Param("path")
	method := c.Method()
	op := blink.NewOp(method, "/api/webdav/broker/"+path, "broker webdav")
	res, err := wdc.hub.Fetch(nil, brkID, op, c.Request().Body)
	if err != nil {
		return err
	}
	ct := res.Header.Get(ship.HeaderContentType)

	return c.Stream(res.StatusCode, ct, res.Body)
}

func (wdc *webdavCtrl) Minion(c *ship.Context) error {
	return nil
}
