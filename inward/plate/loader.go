package plate

import (
	htm "html/template"
	"io"
	"sync"
	"sync/atomic"
	txt "text/template"

	"github.com/vela-ssoc/manager/model"
	"gorm.io/gorm"
)

type tmplRender interface {
	Execute(io.Writer, any) error
}

type tmplLoad struct {
	info  TmplInfo
	done  atomic.Bool
	mutex sync.Mutex
	db    *gorm.DB
	err   error
	rend  tmplRender
}

func (tl *tmplLoad) rendTo(w io.Writer, v any) error {
	if rend, err := tl.parseLoad(); err != nil {
		return err
	} else {
		return rend.Execute(w, v)
	}
}

func (tl *tmplLoad) unset() {
	tl.mutex.Lock()
	tl.done.Store(false)
	tl.mutex.Unlock()
}

func (tl *tmplLoad) parseLoad() (tmplRender, error) {
	if tl.done.Load() {
		return tl.rend, tl.err
	}

	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	if tl.done.Load() {
		return tl.rend, tl.err
	}

	id := tl.info.ID
	plate := &model.Plate{ID: id}
	err := tl.db.Take(&plate).Error
	if err != nil {
		tl.rend = nil
		tl.err = &FindError{ID: id, Err: err}
	} else {
		if plate.Escape {
			tl.rend, tl.err = htm.New(id).Parse(string(plate.Tmpl))
		} else {
			tl.rend, tl.err = txt.New(id).Parse(string(plate.Tmpl))
		}
	}
	tl.done.Store(true)

	return tl.rend, tl.err
}
