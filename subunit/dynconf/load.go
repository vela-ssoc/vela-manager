package dynconf

import (
	"sync"
	"sync/atomic"
)

type eraser interface {
	get() (any, error)
	unset()
}

type findFunc func() (any, error)

func newLoad(fn findFunc) eraser {
	return &safeLoad{find: fn}
}

type safeLoad struct {
	find  findFunc
	mutex sync.Mutex
	done  atomic.Bool
	err   error
	data  any
}

func (sl *safeLoad) get() (any, error) {
	if sl.done.Load() {
		return sl.data, sl.err
	}

	sl.mutex.Lock()
	defer sl.mutex.Lock()

	data, err := sl.find()
	sl.data = data
	sl.err = err
	sl.done.Store(true)

	return data, err
}

func (sl *safeLoad) unset() {
	sl.mutex.Lock()
	sl.done.Store(false)
	sl.mutex.Unlock()
}
