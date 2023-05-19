package mgtapi

import (
	"net/http"
	"strings"
	"time"

	"github.com/vela-ssoc/vela-manager/app/internal/modview"
	"github.com/vela-ssoc/vela-manager/app/internal/param"
	"github.com/vela-ssoc/vela-manager/app/route"
	"github.com/vela-ssoc/vela-manager/app/service"
	"github.com/vela-ssoc/vela-manager/app/session"
	"github.com/xgfone/ship/v5"
)

func Auth(svc service.AuthService) route.Router {
	return &authREST{
		svc: svc,
	}
}

type authREST struct {
	svc service.AuthService
}

// Route 注册路由
func (ath *authREST) Route(anon, bearer, _ *ship.RouteGroupBuilder) {
	anon.Route("/captcha/generate").POST(ath.Picture)
	anon.Route("/captcha/verify").POST(ath.Verify)
	anon.Route("/ding").POST(ath.Dong)
	anon.Route("/login").POST(ath.Login)

	bearer.Route("/logout").DELETE(ath.Logout)
}

// Picture 图片验证码
func (ath *authREST) Picture(c *ship.Context) error {
	var req param.AuthBase
	if err := c.Bind(&req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	res, err := ath.svc.Picture(ctx, req.Username)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

// Verify 验证图片验证码
func (ath *authREST) Verify(c *ship.Context) error {
	var req param.AuthVerify
	if err := c.Bind(&req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	dong, err := ath.svc.Verify(ctx, req)
	if err != nil {
		return err
	}
	res := &param.AuthNeedDong{Ding: dong}

	return c.JSON(http.StatusOK, res)
}

// Dong 获取下发咚咚验证码
func (ath *authREST) Dong(c *ship.Context) error {
	var req param.AuthDong
	if err := c.Bind(&req); err != nil {
		return err
	}
	remoteIP := c.RemoteAddr()
	idx := strings.LastIndex(remoteIP, ":")
	if idx >= 0 {
		remoteIP = remoteIP[:idx]
	}

	ctx := c.Request().Context()
	clientIP := c.ClientIP()
	view := modview.LoginDong{
		Header:   c.Header(),
		RemoteIP: remoteIP,
		ClientIP: clientIP,
		LoginAt:  time.Now(),
	}

	if err := ath.svc.Dong(ctx, req, view); err != nil {
		return err
	}

	return nil
}

// Login 执行登录操作
func (ath *authREST) Login(c *ship.Context) error {
	var req param.AuthLogin
	if err := c.Bind(&req); err != nil {
		return err
	}

	remoteIP := c.RemoteAddr()
	idx := strings.LastIndex(remoteIP, ":")
	if idx >= 0 {
		remoteIP = remoteIP[:idx]
	}

	clientIP := c.ClientIP()
	view := modview.LoginDong{
		Header:   c.Header(),
		RemoteIP: remoteIP,
		ClientIP: clientIP,
		LoginAt:  time.Now(),
	}
	ctx := c.Request().Context()
	user, err := ath.svc.Login(ctx, req, view)
	if err != nil {
		return err
	}

	// 构造 session
	cu := session.Issued(user)
	c.Any = cu
	if err = c.SetSession(cu.Token, cu); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cu)
}

func (ath *authREST) Logout(c *ship.Context) error {
	cu := session.Cast(c.Any)
	return c.DelSession(cu.Token)
}
