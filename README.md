# 安全管理平台

## Go 操作系统支持

2023 年 2 月 1 日，Go 发布了 `Go 1.20` 版本，本项目也第一时间将 Go SDK 升级到了 `Go 1.20`，同时 [Go 1.20 Release Notes](https://go.dev/doc/go1.20) 也告知 `Go 1.20` 将会是最后一个支持 `Windows Server 2008`，`Windows Server 2012` 操作系统的版本。由于当前我们还有运行 `Windows Server 2008`，`Windows Server 2012` 的服务器，为了保证兼容性请 <font color='red'>**谨慎升级 Go SDK 版本**</font>。

> 个人猜测与见解：Go 后续新版本不支持旧的操作系统，不代表不能在旧的操作系统上运行，只是官方不再保证与旧系统的兼容性，如果因为一些原因（如：漏洞）需要升级 Go SDK 的情况，务必做好测试再上线。

### Windows 平台

`Go 1.20` 将会是最后一个支持 `Windows 7，8`，`Windows Server 2008`， `Windows Server 2012` 的版本，`Go1.21` 运行最低的操作系统版本是 `Windows 10` 与 `Windows Server 2016`。

### Darwin(macOS) 与 iOS 平台

`Go 1.20` 将会是最后一个支持 `macOS 10.13 High Sierra` 与 `10.14 Mojave` 的版本，`Go 1.21` 仅支持 `macOS 10.15 Catalina` 或更高版本。
