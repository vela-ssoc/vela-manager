package launch

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/vela-ssoc/backend-common/logback"
	"github.com/vela-ssoc/backend-common/netutil"
	"github.com/vela-ssoc/backend-common/opurl"
	"github.com/vela-ssoc/backend-common/validate"
	"github.com/vela-ssoc/vela-manager/blink"
	"github.com/vela-ssoc/vela-manager/infra/conf"
	"github.com/vela-ssoc/vela-manager/infra/dbms"
	"github.com/vela-ssoc/vela-manager/infra/profile"
	"github.com/vela-ssoc/vela-manager/inward/evtrsk"
	"github.com/vela-ssoc/vela-manager/inward/loadcfg"
	"github.com/vela-ssoc/vela-manager/inward/plate"
	"github.com/vela-ssoc/vela-manager/inward/recisub"
	"github.com/vela-ssoc/vela-manager/outward/sendto"
	"github.com/vela-ssoc/vela-manager/restful"
	"github.com/vela-ssoc/vela-manager/restful/mgtapi"
	"github.com/vela-ssoc/vela-manager/restful/middle"
	"github.com/vela-ssoc/vela-manager/subunit/session"
	"github.com/xgfone/ship/v5"
)

func Run(ctx context.Context, path string) error {
	var cfg conf.Config
	if err := profile.Load(path, &cfg); err != nil {
		return err
	}

	izr := new(initializr)

	return izr.run(ctx, cfg)
}

type initializr struct {
	node  string
	ctx   context.Context
	cfg   conf.Config
	valid validate.Validator
	slog  logback.Logger
	errch chan<- error
}

func (izr *initializr) run(ctx context.Context, cfg conf.Config) error {
	const node = "manager"
	izr.node, izr.ctx, izr.cfg = node, ctx, cfg

	valid := validate.New()
	izr.valid = valid
	if err := valid.Validate(cfg); err != nil { // 校验配置文件
		return err
	}

	zlg := cfg.Logger.Zap()
	db, _, err := dbms.Open(cfg.Database, zlg) // 连接数据库
	if err != nil {
		return err
	}

	slog := logback.Sugar(zlg)
	izr.slog = slog

	// ----------------------[ 开始依赖注入组装 ]----------------------
	hcli := opurl.NewClient()
	rend := plate.DBTmpl(db)                  // 通知模板渲染器
	dongCfg := loadcfg.Dong(db)               // 咚咚服务号配置加载器
	alertCfg := loadcfg.Alert(db)             // 自动化运维告警配置加载器
	dongSend := sendto.Dong(dongCfg, hcli)    // 咚咚通知推送客户端
	alertSend := sendto.Alert(alertCfg, hcli) // 自动化运维告警推送客户端

	dongOpt := sendto.WithDong(dongSend)
	emailOpt := sendto.WithEmail(alertSend)
	wechatOpt := sendto.WithWechat(alertSend)
	smsOpt := sendto.WithSMS(alertSend)
	phoneOpt := sendto.WithPhone(alertSend)
	postman := sendto.Postman(dongOpt, emailOpt, wechatOpt, smsOpt, phoneOpt) // 全局通知推送模块
	subs := recisub.Subscribe(db)                                             // 事件订阅者
	notice := evtrsk.NewHandle(db, subs, rend, postman)                       // 事件/风险通知处理器

	sess := session.DBSess(db, cfg.Server.Sess)
	primary := izr.newShip()
	special := izr.newShip()
	primary.Session = sess
	special.Session = sess

	midOplog := middle.Oplog(nil)
	midAuth := middle.Auth()
	priGroup := primary.Group("/api/v1").Use(midOplog)
	// specGroup := special.Group("/api/v1").Use(midOplog)
	priAnon := priGroup.Clone()
	// specAnon := specGroup.Clone()
	priAuth := priGroup.Use(midAuth)
	// specAuth := specGroup.Use(midAuth)

	// broker 节点请求处理器
	bsh := izr.newShip()
	brk := restful.Brk(bsh, db)
	hub := blink.Hub(db, notice, brk, cfg, slog, node) // broker 节点连接中心
	gw := blink.Gateway(hub)                           // broker 上线连接处理器

	mgtapi.Blink(gw).Route(priAnon, priAuth)
	mgtapi.Into(db, hub, node).Route(priAnon, priAuth)

	// ----------------------[ 结束依赖注入组装 ]----------------------

	h, err := izr.vhostHandler(primary, special)
	if err != nil {
		return err
	}

	_ = hub.ResetDB()

	errch := make(chan error, 1)
	izr.errch = errch
	go izr.listen(h)

	select {
	case <-ctx.Done():
	case err = <-errch:
	}

	return err
}

func (izr *initializr) listen(h http.Handler) {
	svc := izr.cfg.Server
	certs, err := svc.Certs()
	if err != nil {
		izr.errch <- err
		return
	}

	srv := &http.Server{Handler: h, Addr: svc.Addr}
	if len(certs) != 0 {
		srv.TLSConfig = &tls.Config{Certificates: certs}
		// 配置了 TLSConfig 的 Certificates，certFile 和 keyFile 就可以留空了。
		// https://github.com/golang/go/blob/23c0121e4eb259cc1087d0f79a0803cbc71f500b/src/crypto/tls/common.go#L1074-L1107
		err = srv.ListenAndServeTLS("", "")
	} else {
		err = srv.ListenAndServe()
	}
	izr.errch <- err
}

func (izr *initializr) newShip() *ship.Ship {
	const node = "manager"
	sh := ship.Default()
	sh.NotFound = netutil.Notfound(node)
	sh.HandleError = netutil.ErrorFunc(node)
	sh.Validator = izr.valid
	sh.Logger = izr.slog

	return sh
}

func (izr *initializr) vhostHandler(primary, special *ship.Ship) (http.Handler, error) {
	vhosts := izr.cfg.Server.Vhosts
	if len(vhosts) == 0 {
		return primary, nil
	}

	mana := ship.NewHostManagerHandler(nil)
	for _, vhost := range vhosts {
		if _, err := mana.AddHost(vhost, primary); err != nil {
			return nil, err
		}
		mana.SetDefaultHost("", special)
	}

	return mana, nil
}
