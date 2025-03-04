package banner

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

// ANSI 打印 banner 到指定输出流。
//
// 何为 ANSI 转义序列：https://en.wikipedia.org/wiki/ANSI_escape_code.
func ANSI(w io.Writer) {
	parseOnce()
	_, _ = fmt.Fprintf(w, ansiLogo, version, pid, goos, arch,
		hostname, username, workdir, compileAt, commitAt, buildPath, revision)
}

// SupportedANSI 通过客户端 UserAgent 判断是否支持 ANSI 字符。
//
// https://github.com/chubin/wttr.in/blob/355a515b1f9ea31b193cb9c1d3cef5e5dff07de9/internal/processor/processor.go#L24-L41
// https://github.com/chubin/wttr.in/blob/355a515b1f9ea31b193cb9c1d3cef5e5dff07de9/internal/processor/processor.go#L322-L332
func SupportedANSI(userAgent string) bool {
	ua := strings.ToLower(userAgent)
	for _, command := range []string{
		"curl", "httpie", "lwp-request", "wget", "python-httpx", "python-requests",
		"openbsd ftp", "powershell", "fetch", "aiohttp", "http_get", "xh", "nushell",
	} {
		if strings.Contains(ua, command) {
			return true
		}
	}

	return false
}

const ansiLogo = "\033[1;33m" +
	"   ______________  _____  \n" +
	"  / ___/ ___/ __ \\/ ___/ \n" +
	" (__  |__  ) /_/ / /__    \n" +
	"/____/____/\\____/\\___/  \033[0mMANAGER\n" +
	"\033[0;35m:: 东方财富安全管理部 ::\033[0m  \033[3;95m%s\033[0m\n\n" +
	"    \033[1;36m进程 PID:\033[0m %d\n" +
	"    \033[1;36m操作系统:\033[0m %s\n" +
	"    \033[1;36m系统架构:\033[0m %s\n" +
	"    \033[1;36m主机名称:\033[0m %s\n" +
	"    \033[1;36m当前用户:\033[0m %s\n" +
	"    \033[1;36m工作目录:\033[0m %s\n" +
	"    \033[1;36m编译时间:\033[0m %s\n" +
	"    \033[1;36m提交时间:\033[0m %s\n" +
	"    \033[1;36m修订版本:\033[0m https://%s/tree/%s\n\n"

var (
	version     string // 允许 -X 编译时注入
	compileTime string // 允许 -X 编译时注入
	pid         int
	goos        string
	arch        string
	hostname    string
	username    string
	workdir     string
	revision    string
	buildPath   string
	compileAt   time.Time // 处理后的编译时间
	commitAt    time.Time
	parseOnce   = sync.OnceFunc(parse)
)

func parse() {
	pid = os.Getpid()
	goos = runtime.GOOS
	arch = runtime.GOARCH
	hostname, _ = os.Hostname()
	if cu, _ := user.Current(); cu != nil {
		username = cu.Username
	}
	workdir, _ = os.Getwd()
	compileAt = parseTime(compileTime)
	if version == "" && !compileAt.IsZero() {
		version = compileAt.Format("v06.1.2-150405")
	}

	info, _ := debug.ReadBuildInfo()
	if info == nil {
		return
	}
	buildPath = path.Dir(info.Path)
	if version == "" {
		version = info.Main.Version
	}
	settings := info.Settings
	for _, set := range settings {
		key, val := set.Key, set.Value
		switch key {
		case "vcs.revision":
			revision = val
		case "vcs.time":
			commitAt = parseTime(val)
		}
	}
}

func parseTime(str string) time.Time {
	for _, layout := range []string{
		time.RFC1123Z, time.UnixDate, time.Layout, time.ANSIC,
		time.RubyDate, time.RFC822, time.RFC822Z, time.RFC850,
		time.RFC1123, time.RFC3339, time.RFC3339Nano, time.Kitchen,
		time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
		time.DateTime, time.DateOnly,
	} {
		dt, err := time.Parse(layout, str)
		if err != nil || dt.IsZero() {
			continue
		}

		return dt.Local()
	}

	return time.Time{}
}
