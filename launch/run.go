package launch

import (
	"context"
	"net/http"

	"github.com/vela-ssoc/vela-manager/broker"
	"github.com/vela-ssoc/vela-manager/broker/blink"
	"github.com/vela-ssoc/vela-manager/infra/conf"
	"github.com/vela-ssoc/vela-manager/infra/grid"
	"github.com/vela-ssoc/vela-manager/infra/hanerr"
	"github.com/vela-ssoc/vela-manager/infra/logback"
	"github.com/vela-ssoc/vela-manager/inward/evtrsk"
	"github.com/vela-ssoc/vela-manager/inward/loadcfg"
	"github.com/vela-ssoc/vela-manager/inward/plate"
	"github.com/vela-ssoc/vela-manager/inward/recisub"
	"github.com/vela-ssoc/vela-manager/inward/sessm"
	"github.com/vela-ssoc/vela-manager/libkit/httpclient"
	"github.com/vela-ssoc/vela-manager/libkit/profile"
	"github.com/vela-ssoc/vela-manager/libkit/validate"
	"github.com/vela-ssoc/vela-manager/mgtapi"
	"github.com/vela-ssoc/vela-manager/middle"
	"github.com/vela-ssoc/vela-manager/outward/sendto"
	"github.com/xgfone/ship/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Run 项目启动
func Run(parent context.Context, file string) error {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	var cfg conf.Config
	if err := profile.Load(file, &cfg); err != nil { // 加载配置文件
		return err
	}

	// ----------[ 校验配置 ]----------
	valid := validate.New()                     // 参数校验器
	if err := valid.Validate(cfg); err != nil { // 对加载的配置校验
		return err
	}
	// ----------[ 根据配置文件初始化日志 ]----------
	zap := cfg.Logger.Zap()    // 根据配置初始化 zap 日志
	slog := logback.Sugar(zap) // 实例化日志

	// ----------[ 根据配置初始化 gorm 日志并连接数据库 ]----------
	dbCfg := cfg.Database
	glg := logback.GORM(zap, dbCfg.Level) // 初始化 gorm 日志
	dsn := dbCfg.FormatDSN()              // 获取数据库的 DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: glg})
	if err != nil {
		return err
	}
	rawDB, err := db.DB()
	if err != nil {
		return err
	}
	// ----------[ 设置连接参数 ]----------
	rawDB.SetMaxIdleConns(dbCfg.MaxIdleConn)
	rawDB.SetMaxOpenConns(dbCfg.MaxOpenConn)
	rawDB.SetConnMaxLifetime(dbCfg.MaxLifeTime)
	rawDB.SetConnMaxIdleTime(dbCfg.MaxIdleTime)

	gfs := grid.NewCDN(rawDB, dbCfg.CDN, 0)      // 文件存储模块
	httpCli := httpclient.NewClient()            // 创建全局公用的 http client
	rend := plate.DBTmpl(db)                     // 通知模板渲染器
	dongCfg := loadcfg.Dong(db)                  // 咚咚服务号配置加载器
	alertCfg := loadcfg.Alert(db)                // 自动化运维告警配置加载器
	dongSend := sendto.Dong(dongCfg, httpCli)    // 咚咚通知推送客户端
	alertSend := sendto.Alert(alertCfg, httpCli) // 自动化运维告警推送客户端

	dongOpt := sendto.WithDong(dongSend)
	emailOpt := sendto.WithEmail(alertSend)
	wechatOpt := sendto.WithWechat(alertSend)
	smsOpt := sendto.WithSMS(alertSend)
	phoneOpt := sendto.WithPhone(alertSend)
	postman := sendto.Postman(dongOpt, emailOpt, wechatOpt, smsOpt, phoneOpt) // 全局通知推送模块
	subs := recisub.Subscribe(db)                                             // 事件订阅者
	notice := evtrsk.NewHandle(db, subs, rend, postman)                       // 事件/风险通知处理器

	srvCfg := cfg.Server
	sess := sessm.DBSess(db, srvCfg.Sess) // session 管理器

	hostHandler := ship.Default()
	downHandler := ship.Default()

	hostHandler.Session = sess
	downHandler.Session = sess
	hostHandler.Logger = slog
	downHandler.Logger = slog
	hostHandler.Validator = valid
	downHandler.Validator = valid
	hostHandler.HandleError = hanerr.Handle
	downHandler.HandleError = hanerr.Handle
	if dir := srvCfg.HTML; dir != "" {
		// 设置静态代理目录，downHandler 不用设置，
		// 设置 vhost 的目的就是为了防止扫描器直接
		// 扫描 IP 就将我们的网站扫出来
		hostHandler.Route("/").Static(dir)
	}

	midAuth := middle.Auth()
	hostGroup := hostHandler.Group("/api")
	downGroup := downHandler.Group("/api")
	hostAnon := hostGroup.Clone()
	downAnon := downGroup.Clone()
	hostAuth := hostGroup.Use(midAuth)
	downAuth := downGroup.Use(midAuth)

	ping := mgtapi.Ping()
	ping.BindRoute(hostAnon, hostAuth)
	ping.BindRoute(downAnon, downAuth)

	// broker 节点接入相关
	brk := broker.New(db, valid, slog)
	hub := blink.Hub(db, notice, brk, cfg)
	hub.Reset() // 将所有 broker 置为离线状态
	joiner := blink.Gateway(hub)
	link := mgtapi.Blink(joiner)
	link.BindRoute(hostAnon, hostAuth)
	mgtapi.Attach(hub).BindRoute(hostAnon, hostAuth)
	mgtapi.WebDAV("/", hub).BindRoute(hostAnon, hostAuth)

	dep := mgtapi.Deploy(db, gfs)
	dep.BindRoute(hostAnon, hostAuth)
	dep.BindRoute(downAnon, downAuth)

	var handler http.Handler
	if vhosts := srvCfg.Vhosts; len(vhosts) == 0 {
		handler = hostHandler
	} else {
		mana := ship.NewHostManagerHandler(nil)
		for _, host := range vhosts {
			if _, err = mana.AddHost(host, hostHandler); err != nil {
				return err
			}
		}
		mana.SetDefaultHost("", downHandler)
		handler = mana
	}

	errCh := make(chan error)
	daemon := &daemonHTTP{
		config:  srvCfg,
		handler: handler,
		errCh:   errCh,
	}
	go daemon.Run() // 运行 HTTP 服务

	// ----------[ 等待错误信息/结束信号 ]----------
	select {
	case err = <-errCh:
	case <-ctx.Done():
	}

	// ----------[ 程序执行结束关闭资源 ]----------
	_ = daemon.Close() // 关闭 HTTP 服务
	hub.Reset()        // 将所有 broker 置为离线状态
	_ = rawDB.Close()  // 关闭数据库连接
	_ = zap.Sync()     // sync 日志缓冲区

	return err
}
