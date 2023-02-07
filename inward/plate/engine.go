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

// tmplExecutor 抽象出 html 和 text 模板引擎的共同方法
type tmplExecutor interface {
	Execute(io.Writer, any) error
}

// tmplEngine 模板引擎
type tmplEngine struct {
	info  Info         // 模板基本信息
	done  atomic.Bool  // 是否已经加载完毕模板
	mutex sync.Mutex   // 加载模板锁
	db    *gorm.DB     // 数据库连接
	err   error        // 加载或编译模板时产生的错误
	exec  tmplExecutor // 模板执行引擎
}

// unset 模板引擎失效（下次调用会从数据库重新加载最新模板）
func (te *tmplEngine) unset() {
	te.mutex.Lock()
	te.done.Store(false)
	te.mutex.Unlock()
}

// renderTo 渲染模板并输出到指定的 io.Writer 中
func (te *tmplEngine) renderTo(w io.Writer, v any) error {
	exec, err := te.executor()
	if err == nil {
		err = exec.Execute(w, v)
	}
	return err
}

// executor 加载模板引擎
func (te *tmplEngine) executor() (tmplExecutor, error) {
	if te.done.Load() {
		return te.exec, te.err
	}

	te.mutex.Lock()
	defer te.mutex.Unlock()
	if te.done.Load() {
		return te.exec, te.err
	}

	defer te.done.Store(true)

	id := te.info.ID
	pla := &model.Plate{ID: id}
	if err := te.db.Take(pla).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &Error{ID: id}
		}
		te.err = err
		return nil, err
	}

	tpl := string(pla.Tmpl)
	if pla.Escape {
		te.exec, te.err = htm.New(id).Parse(tpl)
	} else {
		te.exec, te.err = txt.New(id).Parse(tpl)
	}

	return te.exec, te.err
}
