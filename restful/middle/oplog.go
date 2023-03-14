package middle

import (
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/vela-ssoc/backend-common/model"
	"github.com/vela-ssoc/vela-manager/restful/facade"
	"github.com/xgfone/ship/v5"
)

// Oplog 操作日志记录中间件，内部包含 recovery
func Oplog(recd facade.Recorder) ship.Middleware {
	newFn := func() any {
		const max = 10 * 1024
		return &limitCopy{
			max:  max,
			data: make([]byte, max),
		}
	}
	m := &oplogMid{
		recd: recd,
		pool: sync.Pool{New: newFn},
	}

	return m.middleware
}

type oplogMid struct {
	pool sync.Pool
	recd facade.Recorder
}

func (m *oplogMid) middleware(handler ship.Handler) ship.Handler {
	return func(c *ship.Context) error {
		body := m.getLimitCopy(c.Body())
		c.Request().Body = body
		begin := time.Now()

		var err error
		defer func() {
			elapsed := time.Since(begin)
			addr, method, uri := c.RemoteAddr(), c.Method(), c.RequestURI()
			if cause := recover(); cause != nil {
				c.WriteHeader(http.StatusInternalServerError)
				c.Errorf("%s %v", elapsed, cause)
			} else {
				c.Debugf("%s %s %s %s", elapsed, addr, method, uri)
			}
			if m.recd == nil {
				return
			}

			data := body.bytes()
			rawURL := c.Request().URL
			oplog := &model.Oplog{
				ClientAddr: c.RemoteAddr(),
				DirectAddr: c.ClientIP(),
				Method:     c.Method(),
				Path:       rawURL.Path,
				Query:      rawURL.RawQuery,
				Length:     c.ContentLength(),
				Content:    data,
				RequestAt:  begin,
				Elapsed:    elapsed,
			}
			if err != nil {
				oplog.Failed, oplog.Cause = true, err.Error()
			} else if code := c.StatusCode(); code >= http.StatusBadRequest {
				oplog.Failed, oplog.Cause = true, strconv.Itoa(code)
			}
			if des, ok := c.Route.Data.(facade.Describer); ok {
				oplog.Name = des.Name()
				oplog.Content = des.Proc(data)
			}
			if info, ok := c.Any.(*model.Userinfo); ok {
				oplog.UserID = info.ID
				oplog.Username = info.Username
				oplog.Nickname = info.Nickname
			}
			_ = m.recd.Record(oplog)
		}()

		err = handler(c)

		return err
	}
}

func (m *oplogMid) getLimitCopy(body io.ReadCloser) *limitCopy {
	lc := m.pool.Get().(*limitCopy)
	lc.body = body
	lc.pos = 0
	return lc
}

func (m *oplogMid) putLimitCopy(lc *limitCopy) {
	m.pool.Put(lc)
}

type limitCopy struct {
	body io.ReadCloser // http body reader
	max  int           // 最大容量
	pos  int           // 偏移量
	data []byte        // 数据 dym
}

func (lc *limitCopy) Read(p []byte) (int, error) {
	n, err := lc.body.Read(p)
	if n > 0 && lc.max > lc.pos {
		num := copy(lc.data[lc.pos:], p)
		lc.pos += num
	}
	return n, err
}

func (lc *limitCopy) Close() error {
	return lc.body.Close()
}

func (lc *limitCopy) bytes() []byte {
	return lc.data[:lc.pos]
}
