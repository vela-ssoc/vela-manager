package conf

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

// Database 数据库配置
type Database struct {
	Level       string            `json:"level"         yaml:"level"`                                  // SQL 日志打印级别
	CDN         string            `json:"cdn"           yaml:"cdn"`                                    // 文件缓存位置
	MaxOpenConn int               `json:"max_open_conn" yaml:"max_open_conn"`                          // 最大连接数
	MaxIdleConn int               `json:"max_idle_conn" yaml:"max_idle_conn"`                          // 最大空闲连接数
	MaxLifeTime time.Duration     `json:"max_life_time" yaml:"max_life_time"`                          // 连接最大存活时长
	MaxIdleTime time.Duration     `json:"max_idle_time" yaml:"max_idle_time"`                          // 空闲连接最大时长
	DSN         string            `json:"dsn"           yaml:"dsn"`                                    // 数据源
	User        string            `json:"user"          yaml:"user"   validate:"required_without=DSN"` // 数据库用户名
	Passwd      string            `json:"passwd"        yaml:"passwd" validate:"required_without=DSN"` // 密码
	Net         string            `json:"net"           yaml:"net"`                                    // 连接协议
	Addr        string            `json:"addr"          yaml:"addr"   validate:"required_without=DSN"` // 连接地址
	DBName      string            `json:"dbname"        yaml:"dbname" validate:"required_without=DSN"` // 库名
	Params      map[string]string `json:"params"        yaml:"params"`                                 // 参数
}

func (db Database) FormatDSN() string {
	if dsn := db.DSN; dsn != "" {
		return dsn
	}

	protocol := db.Net
	if protocol == "" {
		protocol = "tcp"
	}
	cfg := &mysql.Config{
		User:   db.User,
		Passwd: db.Passwd,
		Net:    protocol,
		Addr:   db.Addr,
		DBName: db.DBName,
		Params: db.Params,
	}

	return cfg.FormatDSN()
}
