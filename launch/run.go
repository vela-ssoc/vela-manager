package launch

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/vela-ssoc/manager/broker"
	"github.com/vela-ssoc/manager/broker/blink"
	"github.com/vela-ssoc/manager/infra/conf"
	"github.com/vela-ssoc/manager/inward/dongcfg"
	"github.com/vela-ssoc/manager/inward/evtrsk"
	"github.com/vela-ssoc/manager/inward/plate"
	"github.com/vela-ssoc/manager/inward/recisub"
	"github.com/vela-ssoc/manager/libkit/httpclient"
	"github.com/vela-ssoc/manager/libkit/profile"
	"github.com/vela-ssoc/manager/libkit/validate"
	"github.com/vela-ssoc/manager/mgtapi"
	"github.com/vela-ssoc/manager/middle"
	"github.com/vela-ssoc/manager/outward/sendto"
	"github.com/xgfone/ship/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Run 项目启动
// 本项目各个组件依赖关系尽可能
func Run(parent context.Context, file string) error {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	var cfg conf.Config
	if err := profile.Load(file, &cfg); err != nil { // 加载配置文件
		return err
	}

	valid := validate.New()                     // 参数校验器
	if err := valid.Validate(cfg); err != nil { // 对加载的配置校验
		return err
	}

	// 连接数据库
	dsn := cfg.Database.DSN
	dlg := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dlg})
	if err != nil {
		return err
	}

	httpCli := httpclient.NewClient()                    // 创建全局公用的 http client
	rend := plate.Rend(db)                               // 通知模板渲染器
	dongCfg := dongcfg.Dong(db)                          // 咚咚服务号配置加载器
	dongSend := sendto.Dong(dongCfg, httpCli)            // 咚咚通知推送客户端
	postman := sendto.Postman(sendto.WithDong(dongSend)) // 全局通知推送模块
	subs := recisub.Subscribe(db)                        // 事件订阅者

	sh := ship.Default()
	sh.Validator = valid
	group := sh.Group("/api")
	anon := group.Clone()
	auth := group.Use(middle.Auth()) // 需要认证的路由组

	// broker 节点接入相关
	brk := broker.New(db, valid, nil)

	notice := evtrsk.NewHandle(db, subs, rend, postman)
	hub := blink.Hub(db, notice, brk)
	hub.Reset() // 将所有 broker 置为离线状态
	joiner := blink.Gateway(hub)
	mgtapi.Blink(joiner).BindRoute(anon, auth)

	errCh := make(chan error)
	daemon := &daemonHTTP{
		config:  cfg.Server,
		handler: sh,
		errCh:   errCh,
	}
	go daemon.Run() // 运行 HTTP 服务

	select {
	case err = <-errCh:
	case <-ctx.Done():
		err = ctx.Err()
	}
	_ = daemon.Close()

	hub.Reset() // 将所有 broker 置为离线状态

	// 关闭数据库连接
	if sdb, _ := db.DB(); sdb != nil {
		_ = sdb.Close()
	}

	return err
}
