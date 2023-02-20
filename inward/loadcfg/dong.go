package loadcfg

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/vela-ssoc/backend-common/model"
	"github.com/vela-ssoc/vela-manager/outward/sendto"
	"gorm.io/gorm"
)

var ErrNotFoundDong = errors.New("没有找到咚咚服务号配置")

func Dong(db *gorm.DB) sendto.DongConfigurer {
	return &dongConfigure{db: db}
}

type dongConfigure struct {
	db    *gorm.DB
	mutex sync.Mutex
	done  atomic.Bool
	err   error
	data  *model.Dong
}

func (dc *dongConfigure) DongUnset() {
	dc.mutex.Lock()
	dc.done.Store(false)
	dc.mutex.Unlock()
}

func (dc *dongConfigure) DongConfig() (*model.Dong, error) {
	if dc.done.Load() {
		return dc.data, dc.err
	}

	dc.mutex.Lock()
	defer dc.mutex.Unlock()

	if dc.done.Load() {
		return dc.data, dc.err
	}

	data := new(model.Dong)
	err := dc.db.First(data, "enable = ?", true).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrNotFoundDong
	}

	dc.data, dc.err = data, err
	dc.done.Store(true)

	return data, err
}
