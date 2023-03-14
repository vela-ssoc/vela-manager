package blink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

// Joiner broker 节点接入
type Joiner interface {
	// Auth 开始认证
	Auth(Ident) (Issue, http.Header, error)
	// Join 认证成功后接入处理业务逻辑
	Join(net.Conn, Ident, Issue) error
}

// Gateway 上线接入网关
func Gateway(joiner Joiner) http.Handler {
	return &gateway{joiner: joiner}
}

type gateway struct {
	joiner Joiner
}

// ServeHTTP 请求接入
func (gw *gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 验证 HTTP 方法
	if method := r.Method; method != http.MethodConnect {
		gw.writeError(w, http.StatusBadRequest, "不支持的请求方法：%s", method)
		return
	}

	buf := make([]byte, 100*1024)
	n, _ := io.ReadFull(r.Body, buf)
	var ident Ident
	if err := ident.decrypt(buf[:n]); err != nil {
		gw.writeError(w, http.StatusBadRequest, "认证信息错误")
		return
	}

	// 鉴权
	issue, header, gex := gw.joiner.Auth(ident)
	if gex != nil {
		gw.writeError(w, http.StatusBadRequest, "认证失败：%s", gex.Error())
		return
	}

	dat, err := issue.encrypt()
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

	if err = gw.joiner.Join(conn, ident, issue); err != nil {
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
