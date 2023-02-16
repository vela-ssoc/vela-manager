package conf

import (
	"crypto/tls"
	"time"
)

type Server struct {
	Addr   string        `json:"addr"   yaml:"addr"`
	Cert   string        `json:"cert"   yaml:"cert"`
	Pkey   string        `json:"pkey"   yaml:"pkey"`
	HTML   string        `json:"html"   yaml:"html"`
	Sess   time.Duration `json:"sess"   yaml:"sess"`
	Vhosts []string      `json:"vhosts" yaml:"vhosts" validate:"dive,required"`
}

func (srv Server) Certs() ([]tls.Certificate, error) {
	if srv.Cert == "" || srv.Pkey == "" {
		return nil, nil
	}

	cert, err := tls.LoadX509KeyPair(srv.Cert, srv.Pkey)
	if err != nil {
		return nil, err
	}
	certs := []tls.Certificate{cert}

	return certs, nil
}
