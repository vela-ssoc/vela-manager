package loadcfg

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/vela-ssoc/backend-common/model"
	"github.com/vela-ssoc/vela-manager/outward/sendto"
	"gorm.io/gorm"
)

var ErrNotFoundAlert = errors.New("没有找到自动化运维告警配置")

func Alert(db *gorm.DB) sendto.AlertConfigurer {
	return &alertConfigure{db: db}
}

type alertConfigure struct {
	db    *gorm.DB
	mutex sync.Mutex
	done  atomic.Bool
	err   error
	data  *model.Alert
}

func (ac *alertConfigure) AlertUnset() {
	ac.mutex.Lock()
	ac.done.Store(false)
	ac.mutex.Unlock()
}

func (ac *alertConfigure) AlertConfig() (*model.Alert, error) {
	if ac.done.Load() {
		return ac.data, ac.err
	}

	ac.mutex.Lock()
	defer ac.mutex.Unlock()

	if ac.done.Load() {
		return ac.data, ac.err
	}

	data := new(model.Alert)
	err := ac.db.First(data, "enable = ?", true).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrNotFoundAlert
	}

	ac.data, ac.err = data, err
	ac.done.Store(true)

	return data, err
}
