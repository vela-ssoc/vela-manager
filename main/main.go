package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vela-ssoc/vela-manager/infra/banner"
	"github.com/vela-ssoc/vela-manager/launch"
)

// main 安全平台管理端
func main() {
	// 处理命令行输入的参数
	ver := flag.Bool("v", false, "打印版本号并退出")
	cfg := flag.String("c", "zone/conf/manager.yaml", "配置文件路径")
	flag.Parse()

	if banner.Print(os.Stdout); *ver {
		return
	}

	// 监听停止信号
	cares := []os.Signal{syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGINT}
	ctx, cancel := signal.NotifyContext(context.Background(), cares...)
	defer cancel()
	log.Println("按 Ctrl+C 结束运行")

	if err := launch.Run(ctx, *cfg); err != nil /*&& err != context.Canceled*/ {
		log.Printf("程序启动错误：%s", err)
	}

	log.Println("程序已停止运行")
}
