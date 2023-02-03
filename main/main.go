package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vela-ssoc/manager/infra/banner"
	"github.com/vela-ssoc/manager/launch"
)

func main() {
	// 处理命令行输入的参数
	var config string
	var version bool
	flag.StringVar(&config, "c", "zone/conf/manager.yaml", "配置文件")
	flag.BoolVar(&version, "v", false, "打印版本号")
	flag.Parse()

	if banner.Print(); version {
		return
	}

	// 监听停止信号
	cares := []os.Signal{syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGINT}
	ctx, cancel := signal.NotifyContext(context.Background(), cares...)
	defer cancel()
	log.Println("按 Ctrl+C 结束运行")

	if err := launch.Run(ctx, config); err != nil && err != context.Canceled {
		log.Println(err)
	}

	log.Println("程序已停止运行")
}
