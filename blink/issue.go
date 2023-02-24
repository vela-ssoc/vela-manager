package blink

import (
	"encoding/json"

	"github.com/vela-ssoc/backend-common/encipher"
	"github.com/vela-ssoc/vela-manager/infra/conf"
)

// Issue broker 节点认证成功后返回的信息
type Issue struct {
	Name     string        `json:"name"`     // broker 名字
	Passwd   []byte        `json:"passwd"`   // 通信加解密密钥
	Listen   Listen        `json:"listen"`   // 服务监听配置
	Logger   conf.Logger   `json:"logger"`   // 日志配置
	Database conf.Database `json:"database"` // 数据库配置
}

// Listen 监听信息
type Listen struct {
	Addr string `json:"addr"` // 监听地址 :8080 192.168.1.2:8080
	Cert []byte `json:"cert"` // 证书
	Pkey []byte `json:"pkey"` // 私钥
}

// String fmt.Stringer
func (iss Issue) String() string {
	dat, _ := json.MarshalIndent(iss, "", "    ")
	return string(dat)
}

func (iss Issue) Encrypt() ([]byte, error) {
	return encipher.EncryptJSON(iss)
}
