package dynconf

import (
	"errors"

	"github.com/vela-ssoc/backend-common/model"
	"gorm.io/gorm"
)

var ErrNotfoundDong = errors.New("没有找到咚咚服务号配置")

func Dong(db *gorm.DB) {
	df := &dongFind{db: db}
}

type dongFind struct {
	db    *gorm.DB
	erase eraser
}

func (df *dongFind) find() (any, error) {
	data := new(model.Dong)
	err := df.db.First(data, "enable = ?", true).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrNotfoundDong
	}
	return data, err
}
