# 安全管理平台

## 中心端开发部署

> 本项目使用 go.work 方式组织项目： [multi-module workspaces](https://go.dev/doc/tutorial/workspaces)

```shell
# 创建工作目录
mkdir vela-ssoc

# 进入工作目录，拉取代码
# backend-common: 后端公共代码库
# vela-manager: 中心端管理端
# vela-broker: 代理节点
# vela-minion: agent 节点
cd vela-ssoc
git clone https://github.com/vela-ssoc/backend-common.git
git clone https://github.com/vela-ssoc/vela-manager.git
git clone https://github.com/vela-ssoc/vela-broker.git
git clone https://github.com/vela-ssoc/vela-minion.git

# 初始化 go.work
go work init
go work use ./backend-common
go work use ./vela-manager
go work use ./vela-broker
go work use ./vela-minion

# GoLand/IDEA 打开 ssoc-central 目录即可开发 manager 和 broker 模块
```

## Go SDK 升级注意事项

2023 年 2 月 1 日，Go 发布了 `Go 1.20` 版本，本项目也第一时间将 Go SDK 升级到了 `Go 1.20`，同时 [Go 1.20 Release Notes](https://go.dev/doc/go1.20) 也告知 `Go 1.20` 将会是最后一个支持 `Windows Server 2008`，`Windows Server 2012` 操作系统的版本。由于当前我们还有运行 `Windows Server 2008`，`Windows Server 2012` 的服务器，为了保证兼容性请 <font color='red'>**谨慎升级 Go SDK 版本**</font>。

> 个人猜测与见解：Go 后续新版本不支持旧的操作系统，不代表不能在旧的操作系统上运行，只是官方不再保证与旧系统的兼容性，如果因为一些原因（如：漏洞）需要升级 Go SDK 的情况，务必做好测试再上线。

### Windows 平台

`Go 1.20` 将会是最后一个支持 `Windows 7，8`，`Windows Server 2008`， `Windows Server 2012` 的版本，`Go1.21` 运行最低的操作系统版本是 `Windows 10` 与 `Windows Server 2016`。

### Darwin(macOS) 与 iOS 平台

`Go 1.20` 将会是最后一个支持 `macOS 10.13 High Sierra` 与 `10.14 Mojave` 的版本，`Go 1.21` 仅支持 `macOS 10.15 Catalina` 或更高版本。
