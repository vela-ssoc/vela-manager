package conf

import "crypto/tls"

type Config struct {
	Server   Server   `json:"server"   yaml:"server"`
	Database Database `json:"database" yaml:"database"`
}

type Server struct {
	Addr string `json:"addr" yaml:"addr"`
	Cert string `json:"cert" yaml:"cert"`
	Pkey string `json:"pkey" yaml:"pkey"`
}

type Database struct {
	DSN string `json:"dsn" yaml:"dsn"`
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
