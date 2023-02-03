package plate

import (
	"bytes"
	"fmt"
	"io"

	"gorm.io/gorm"
)

const (
	tplEventDongTitle       = "event.dong.title"
	tplEventDongContent     = "event.dong.content"
	tplEventEmailTitle      = "event.email.title"
	tplEventEmailContent    = "event.email.content"
	tplEventWechatTitle     = "event.wechat.title"
	tplEventWechatContent   = "event.wechat.content"
	tplEventSMSContent      = "event.sms.content"
	tplEventPhoneContent    = "event.phone.content"
	tplRiskDongTitle        = "risk.dong.title"
	tplRiskDongContent      = "risk.dong.content"
	tplRiskEmailTitle       = "risk.email.title"
	tplRiskEmailContent     = "risk.email.content"
	tplRiskWechatTitle      = "risk.wechat.title"
	tplRiskWechatContent    = "risk.wechat.content"
	tplRiskSMSContent       = "risk.sms.content"
	tplRiskPhoneContent     = "risk.phone.content"
	tplLoginCodeDongTitle   = "login.code.dong.title"
	tplLoginCodeDongContent = "login.code.dong.content"
	tplStartupContent       = "startup.content"
)

type Render interface {
	Unset(string)
	Shared(string) bool
	Rightful(string) bool
	Support() []TmplInfo
	Rend(string, any) (*bytes.Buffer, error)
	RendTo(io.Writer, string, any) error
	EventDongTitle(any) (*bytes.Buffer, error)
	EventDongContent(any) (*bytes.Buffer, error)
	EventEmailTitle(any) (*bytes.Buffer, error)
	EventEmailContent(any) (*bytes.Buffer, error)
	EventWechatTitle(any) (*bytes.Buffer, error)
	EventWechatContent(any) (*bytes.Buffer, error)
	EventSMSContent(any) (*bytes.Buffer, error)
	EventPhoneContent(any) (*bytes.Buffer, error)
	RiskDongTitle(any) (*bytes.Buffer, error)
	RiskDongContent(any) (*bytes.Buffer, error)
	RiskEmailTitle(any) (*bytes.Buffer, error)
	RiskEmailContent(any) (*bytes.Buffer, error)
	RiskWechatTitle(any) (*bytes.Buffer, error)
	RiskWechatContent(any) (*bytes.Buffer, error)
	RiskSMSContent(any) (*bytes.Buffer, error)
	RiskPhoneContent(any) (*bytes.Buffer, error)
	LoginCodeDongTitle(any) (*bytes.Buffer, error)
	LoginCodeDongContent(any) (*bytes.Buffer, error)
	StartupContent(any) (*bytes.Buffer, error)
}

type TmplInfo struct {
	ID    string `json:"id"`
	Share bool   `json:"share"`
	Desc  string `json:"desc"`
}

func Rend(db *gorm.DB) Render {
	sups := []TmplInfo{
		{ID: tplEventDongTitle, Share: true, Desc: ""},
		{ID: tplEventDongContent, Share: true, Desc: ""},
		{ID: tplEventEmailTitle, Share: true, Desc: ""},
		{ID: tplEventEmailContent, Share: true, Desc: ""},
		{ID: tplEventWechatTitle, Share: true, Desc: ""},
		{ID: tplEventWechatContent, Share: true, Desc: ""},
		{ID: tplEventSMSContent, Share: true, Desc: ""},
		{ID: tplEventPhoneContent, Share: true, Desc: ""},
		{ID: tplRiskDongTitle, Share: true, Desc: ""},
		{ID: tplRiskDongContent, Share: true, Desc: ""},
		{ID: tplRiskEmailTitle, Share: true, Desc: ""},
		{ID: tplRiskEmailContent, Share: true, Desc: ""},
		{ID: tplRiskWechatTitle, Share: true, Desc: ""},
		{ID: tplRiskWechatContent, Share: true, Desc: ""},
		{ID: tplRiskSMSContent, Share: true, Desc: ""},
		{ID: tplRiskPhoneContent, Share: true, Desc: ""},
		{ID: tplLoginCodeDongTitle, Desc: ""},
		{ID: tplLoginCodeDongContent, Desc: ""},
		{ID: tplStartupContent, Desc: ""},
	}

	rends := make(map[string]*tmplLoad, len(sups))
	for _, info := range sups {
		id := info.ID
		rends[id] = &tmplLoad{info: info, db: db}
	}

	return &tmplManager{support: sups, rends: rends}
}

type tmplManager struct {
	support []TmplInfo
	rends   map[string]*tmplLoad
}

func (tm *tmplManager) Unset(id string) {
	if rend, ok := tm.rends[id]; ok {
		rend.unset()
	}
}

func (tm *tmplManager) Shared(id string) bool {
	rend, ok := tm.rends[id]
	return ok && rend.info.Share
}

func (tm *tmplManager) Rightful(id string) bool {
	_, ok := tm.rends[id]
	return ok
}

func (tm *tmplManager) Support() []TmplInfo {
	return tm.support
}

func (tm *tmplManager) Rend(id string, v any) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := tm.RendTo(buf, id, v)
	return buf, err
}

func (tm *tmplManager) RendTo(w io.Writer, id string, v any) error {
	if rend, ok := tm.rends[id]; ok {
		return rend.rendTo(w, v)
	}
	return &FindError{ID: id}
}

func (tm *tmplManager) EventDongTitle(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplEventDongTitle, v)
}

func (tm *tmplManager) EventDongContent(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplEventDongContent, v)
}

func (tm *tmplManager) EventEmailTitle(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplEventEmailTitle, v)
}

func (tm *tmplManager) EventEmailContent(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplEventEmailContent, v)
}

func (tm *tmplManager) EventWechatTitle(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplEventWechatTitle, v)
}

func (tm *tmplManager) EventWechatContent(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplEventWechatContent, v)
}

func (tm *tmplManager) EventSMSContent(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplEventSMSContent, v)
}

func (tm *tmplManager) EventPhoneContent(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplEventPhoneContent, v)
}

func (tm *tmplManager) RiskDongTitle(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplRiskDongTitle, v)
}

func (tm *tmplManager) RiskDongContent(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplRiskDongContent, v)
}

func (tm *tmplManager) RiskEmailTitle(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplRiskEmailTitle, v)
}

func (tm *tmplManager) RiskEmailContent(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplRiskEmailContent, v)
}

func (tm *tmplManager) RiskWechatTitle(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplRiskWechatTitle, v)
}

func (tm *tmplManager) RiskWechatContent(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplRiskWechatContent, v)
}

func (tm *tmplManager) RiskSMSContent(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplRiskSMSContent, v)
}

func (tm *tmplManager) RiskPhoneContent(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplRiskPhoneContent, v)
}

func (tm *tmplManager) LoginCodeDongTitle(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplLoginCodeDongTitle, v)
}

func (tm *tmplManager) LoginCodeDongContent(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplLoginCodeDongContent, v)
}

func (tm *tmplManager) StartupContent(v any) (*bytes.Buffer, error) {
	return tm.Rend(tplStartupContent, v)
}

type FindError struct {
	ID  string
	Err error
}

func (fe *FindError) Error() string {
	if fe.Err == nil {
		return fmt.Sprintf("使用不存在的模板: %s", fe.ID)
	}
	if fe.Err == gorm.ErrRecordNotFound {
		return fmt.Sprintf("没有配置 %s 模板", fe.ID)
	}
	return fmt.Sprintf("查找模板 %s 错误: %s", fe.ID, fe.Err.Error())
}
