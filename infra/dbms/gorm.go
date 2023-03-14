package dbms

import (
	"database/sql"

	"github.com/vela-ssoc/backend-common/logback"
	"github.com/vela-ssoc/vela-manager/infra/conf"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open(cfg conf.Database, zlg *zap.Logger) (*gorm.DB, *sql.DB, error) {
	dsn := cfg.FormatDSN()
	lg := logback.Gorm(zlg, cfg.Level)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: lg})
	if err != nil {
		return nil, nil, err
	}
	sdb, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	// ----------[ 设置连接参数 ]----------
	sdb.SetMaxIdleConns(cfg.MaxIdleConn)
	sdb.SetMaxOpenConns(cfg.MaxOpenConn)
	sdb.SetConnMaxLifetime(cfg.MaxLifeTime)
	sdb.SetConnMaxIdleTime(cfg.MaxIdleTime)

	sn := newSnow()
	_ = db.Callback().
		Create().
		Before("gorm:create").
		Register("generate_id", sn.plugin)

	return db, sdb, nil
}
