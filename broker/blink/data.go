package blink

import "github.com/vela-ssoc/manager/infra/conf"

type Issue struct {
	Name     string        `json:"name"`
	Passwd   []byte        `json:"passwd"`
	Listen   Listen        `json:"listen"`   // 服务监听配置
	Logger   conf.Logger   `json:"logger"`   // 日志配置
	Database conf.Database `json:"database"` // 数据库配置
}

type Listen struct {
	Addr string `json:"addr"` // 监听地址 :8080 192.168.1.2:8080
	Cert []byte `json:"cert"` // 证书
	Pkey []byte `json:"pkey"` // 私钥
}
