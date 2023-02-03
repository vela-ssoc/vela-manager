package blink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/vela-ssoc/manager/infra/encipher"
)

type Ident struct {
	ID         int64     `json:"id"`         // ID
	Secret     string    `json:"secret"`     // 密钥
	IP         net.IP    `json:"ip"`         // IP 地址
	MAC        string    `json:"mac"`        // MAC 地址
	Semver     string    `json:"semver"`     // 版本
	Goos       string    `json:"goos"`       // runtime.GOOS
	Arch       string    `json:"arch"`       // runtime.GOARCH
	CPU        int       `json:"cpu"`        // runtime.NumCPU
	PID        int       `json:"pid"`        // os.Getpid
	Workdir    string    `json:"workdir"`    // os.Getwd
	Executable string    `json:"executable"` // os.Executable
	Username   string    `json:"username"`   // user.Current
	Hostname   string    `json:"hostname"`   // os.Hostname
	TimeAt     time.Time `json:"time_at"`    // 发起时间
}

type Joiner interface {
	Auth(ident Ident) (any, http.Header, error)
	Join(net.Conn, Ident, any) error
}

func Gateway(joiner Joiner) http.Handler {
	return &gateway{joiner: joiner}
}

type gateway struct {
	joiner Joiner
}

func (gw *gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 验证 HTTP 方法
	if method := r.Method; method != http.MethodConnect {
		gw.writeError(w, http.StatusBadRequest, "不支持的请求方法：%s", method)
		return
	}

	buf := make([]byte, 40960)
	n, _ := io.ReadFull(r.Body, buf)
	var ident Ident
	if err := encipher.DecryptJSON(buf[:n], &ident); err != nil {
		gw.writeError(w, http.StatusBadRequest, "认证信息错误")
		return
	}

	// 鉴权
	grant, header, gex := gw.joiner.Auth(ident)
	if gex != nil {
		gw.writeError(w, http.StatusBadRequest, "认证失败：%s", gex.Error())
		return
	}

	dat, err := json.Marshal(grant)
	if err != nil {
		gw.writeError(w, http.StatusInternalServerError, "内部错误：%s", err.Error())
		return
	}

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		gw.writeError(w, http.StatusBadRequest, "协议错误")
		return
	}
	conn, _, jex := hijacker.Hijack()
	if jex != nil {
		gw.writeError(w, http.StatusBadRequest, "协议升级失败：%s", jex.Error())
		return
	}

	// -----[ Hijack Successful ]-----

	// 默认规定 http.StatusAccepted 为成功状态码
	code := http.StatusAccepted
	res := &http.Response{
		Status:     http.StatusText(code),
		StatusCode: code,
		Proto:      r.Proto,
		ProtoMajor: r.ProtoMajor,
		ProtoMinor: r.ProtoMinor,
		Header:     header,
		Request:    r,
	}
	if dsz := len(dat); dsz > 0 {
		res.Body = io.NopCloser(bytes.NewReader(dat))
		res.ContentLength = int64(dsz)
	}
	if err = res.Write(conn); err != nil {
		_ = conn.Close()
		return
	}

	if err = gw.joiner.Join(conn, ident, grant); err != nil {
		_ = conn.Close()
	}
}

// writeError 写入错误
func (*gateway) writeError(w http.ResponseWriter, code int, msg string, args ...string) {
	if code < http.StatusBadRequest || code > http.StatusNetworkAuthenticationRequired {
		code = http.StatusBadRequest
	}

	if len(args) != 0 {
		msg = fmt.Sprintf(msg, args)
	}
	ret := struct {
		Message string `json:"message"`
	}{Message: msg}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(ret)
}
