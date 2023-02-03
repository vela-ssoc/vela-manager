package launch

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"

	"github.com/vela-ssoc/manager/infra/conf"
)

type daemonHTTP struct {
	config  conf.Server  // HTTP 服务启动配置
	handler http.Handler // HTTP Handler
	errCh   chan<- error // 错误 chan
	server  *http.Server // HTTP server
}

func (dh *daemonHTTP) Run() {
	certs, err := dh.config.Certs()
	if err != nil {
		dh.errCh <- err
		return
	}

	srv := &http.Server{
		Addr:     dh.config.Addr,
		Handler:  dh.handler,
		ErrorLog: log.New(io.Discard, "", log.LstdFlags),
	}
	dh.server = srv

	if len(certs) == 0 { // 未启用 TLS
		err = srv.ListenAndServe()
	} else {
		// 配置好 TLSConfig 后，srv.ListenAndServeTLS 就无需填写 certFile 和 keyFile，
		// TLS 在握手时证书匹配规则请参考：
		// https://github.com/golang/go/blob/23c0121e4eb259cc1087d0f79a0803cbc71f500b/src/crypto/tls/common.go#L1074-L1107
		srv.TLSConfig = &tls.Config{Certificates: certs}
		err = srv.ListenAndServeTLS("", "") // 此处留空则代表
	}

	dh.errCh <- err
}

func (dh *daemonHTTP) Close() error {
	if srv := dh.server; srv != nil {
		return srv.Close()
	}
	return nil
}
