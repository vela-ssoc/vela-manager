package mgtapi

import (
	"net/http"
	"strings"
	"time"

	"github.com/vela-ssoc/ssoc-manager/app/route"
	"github.com/vela-ssoc/ssoc-manager/app/service"
	"github.com/vela-ssoc/ssoc-manager/app/session"
	"github.com/vela-ssoc/ssoc-manager/param/modview"
	"github.com/vela-ssoc/ssoc-manager/param/mrequest"
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
	// anon.Route("/captcha/generate").Data(route.Ignore()).POST(ath.Picture)
	// anon.Route("/captcha/verify").Data(route.Ignore()).POST(ath.Verify)
	// anon.Route("/ding").Data(route.Ignore()).POST(ath.Dong)
	// anon.Route("/login").Data(route.DestPasswd("用户登录")).POST(ath.Login)

	anon.Route("/auth/valid").Data(route.DestPasswd("校验用户名密码")).POST(ath.Valid)
	anon.Route("/auth/totp").Data(route.Named("获取 TOTP")).POST(ath.Totp)
	anon.Route("/auth/submit").Data(route.DestPasswd("用户登录")).POST(ath.Submit)
	anon.Route("/auth/oauth").Data(route.DestPasswd("用户登录(oauth)")).POST(ath.Oauth)

	bearer.Route("/logout").Data(route.Named("用户退出登录")).DELETE(ath.Logout)
}

// Picture 图片验证码
func (ath *authREST) Picture(c *ship.Context) error {
	var req mrequest.AuthBase
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
	var req mrequest.AuthVerify
	if err := c.Bind(&req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	ding, err := ath.svc.Verify(ctx, req)
	if err != nil {
		return err
	}
	res := &mrequest.AuthNeedDong{Ding: ding}

	return c.JSON(http.StatusOK, res)
}

// Dong 获取下发咚咚验证码
func (ath *authREST) Dong(c *ship.Context) error {
	var req mrequest.AuthDong
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

	return ath.svc.Dong(ctx, req, view)
}

// Login 执行登录操作
func (ath *authREST) Login(c *ship.Context) error {
	var req mrequest.AuthLogin
	if err := c.Bind(&req); err != nil {
		return err
	}

	remoteIP := c.RemoteAddr()
	idx := strings.LastIndex(remoteIP, ":")
	if idx >= 0 {
		remoteIP = remoteIP[:idx]
	}

	ctx := c.Request().Context()
	user, err := ath.svc.Login(ctx, req)
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

func (ath *authREST) Valid(c *ship.Context) error {
	var req mrequest.AuthBase
	if err := c.Bind(&req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	uid, bind, err := ath.svc.Valid(ctx, req.Username, req.Password)
	if err != nil {
		return err
	}

	res := &mrequest.AuthValidResp{
		UID:  uid,
		Bind: bind,
	}

	return c.JSON(http.StatusOK, res)
}

func (ath *authREST) Totp(c *ship.Context) error {
	var req mrequest.AuthUID
	if err := c.Bind(&req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	otp, err := ath.svc.Totp(ctx, req.UID)
	if err != nil {
		return err
	}

	res := &mrequest.AuthTotpResp{
		TOTP: otp,
		URL:  otp.String(),
	}

	return c.JSON(http.StatusOK, res)
}

func (ath *authREST) Submit(c *ship.Context) error {
	var req mrequest.AuthSubmit
	if err := c.Bind(&req); err != nil {
		return err
	}

	// 查询 UID 是否有效
	ctx := c.Request().Context()
	user, err := ath.svc.Submit(ctx, req.UID, req.Code)
	if err != nil {
		return err
	}

	cu := session.Issued(user)
	c.Any = cu
	if err = c.SetSession(cu.Token, cu); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cu)
}

// Oauth 通过咚咚扫码登录。
func (ath *authREST) Oauth(c *ship.Context) error {
	req := new(mrequest.AuthOauth)
	if err := c.Bind(req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	user, err := ath.svc.Oauth(ctx, req)
	if err != nil {
		return err
	}

	cu := session.Issued(user)
	c.Any = cu
	if err = c.SetSession(cu.Token, cu); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cu)
}
