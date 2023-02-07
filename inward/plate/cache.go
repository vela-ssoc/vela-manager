package plate

import (
	"bytes"
	"io"

	"gorm.io/gorm"
)

type Error struct {
	ID string
}

func (e *Error) Error() string {
	return "模板不存在: " + e.ID
}

type Info struct {
	ID    string
	Desc  string
	Share bool
}

// Render 模板渲染接口规范
type Render interface {
	// Unset 丢弃已经加载的模板，下载使用重新加载最新模板
	Unset(string)
	Shared(string) bool
	Rightful(string) bool
	Infos() []Info
	Render(string, any) (*bytes.Buffer, error)
	RenderTo(io.Writer, string, any) error
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

// DBTmpl 存放在数据库的模板
func DBTmpl(db *gorm.DB) Render {
	infos := []Info{
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

	engines := make(map[string]*tmplEngine, len(infos))
	for _, info := range infos {
		id := info.ID
		engines[id] = &tmplEngine{info: info, db: db}
	}

	return &cachedTmpl{
		infos:   infos,
		engines: engines,
	}
}

type cachedTmpl struct {
	infos   []Info
	engines map[string]*tmplEngine
}

func (ct *cachedTmpl) Unset(id string) {
	if eng := ct.engines[id]; eng != nil {
		eng.unset()
	}
}

func (ct *cachedTmpl) Shared(id string) bool {
	if eng := ct.engines[id]; eng != nil {
		return eng.info.Share
	}
	return false
}

func (ct *cachedTmpl) Rightful(id string) bool {
	_, ok := ct.engines[id]
	return ok
}

func (ct *cachedTmpl) Infos() []Info {
	return ct.infos
}

func (ct *cachedTmpl) Render(id string, v any) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	if err := ct.RenderTo(buf, id, v); err != nil {
		return nil, err
	}
	return buf, nil
}

func (ct *cachedTmpl) RenderTo(w io.Writer, id string, v any) error {
	if eng := ct.engines[id]; eng != nil {
		return eng.renderTo(w, v)
	}
	return &Error{ID: id}
}

func (ct *cachedTmpl) EventDongTitle(v any) (*bytes.Buffer, error) {
	return ct.Render(tplEventDongTitle, v)
}

func (ct *cachedTmpl) EventDongContent(v any) (*bytes.Buffer, error) {
	return ct.Render(tplEventDongContent, v)
}

func (ct *cachedTmpl) EventEmailTitle(v any) (*bytes.Buffer, error) {
	return ct.Render(tplEventEmailTitle, v)
}

func (ct *cachedTmpl) EventEmailContent(v any) (*bytes.Buffer, error) {
	return ct.Render(tplEventEmailContent, v)
}

func (ct *cachedTmpl) EventWechatTitle(v any) (*bytes.Buffer, error) {
	return ct.Render(tplEventWechatTitle, v)
}

func (ct *cachedTmpl) EventWechatContent(v any) (*bytes.Buffer, error) {
	return ct.Render(tplEventWechatContent, v)
}

func (ct *cachedTmpl) EventSMSContent(v any) (*bytes.Buffer, error) {
	return ct.Render(tplEventSMSContent, v)
}

func (ct *cachedTmpl) EventPhoneContent(v any) (*bytes.Buffer, error) {
	return ct.Render(tplEventPhoneContent, v)
}

func (ct *cachedTmpl) RiskDongTitle(v any) (*bytes.Buffer, error) {
	return ct.Render(tplRiskDongTitle, v)
}

func (ct *cachedTmpl) RiskDongContent(v any) (*bytes.Buffer, error) {
	return ct.Render(tplRiskDongContent, v)
}

func (ct *cachedTmpl) RiskEmailTitle(v any) (*bytes.Buffer, error) {
	return ct.Render(tplRiskEmailTitle, v)
}

func (ct *cachedTmpl) RiskEmailContent(v any) (*bytes.Buffer, error) {
	return ct.Render(tplRiskEmailContent, v)
}

func (ct *cachedTmpl) RiskWechatTitle(v any) (*bytes.Buffer, error) {
	return ct.Render(tplRiskWechatTitle, v)
}

func (ct *cachedTmpl) RiskWechatContent(v any) (*bytes.Buffer, error) {
	return ct.Render(tplRiskWechatContent, v)
}

func (ct *cachedTmpl) RiskSMSContent(v any) (*bytes.Buffer, error) {
	return ct.Render(tplRiskSMSContent, v)
}

func (ct *cachedTmpl) RiskPhoneContent(v any) (*bytes.Buffer, error) {
	return ct.Render(tplRiskPhoneContent, v)
}

func (ct *cachedTmpl) LoginCodeDongTitle(v any) (*bytes.Buffer, error) {
	return ct.Render(tplLoginCodeDongTitle, v)
}

func (ct *cachedTmpl) LoginCodeDongContent(v any) (*bytes.Buffer, error) {
	return ct.Render(tplLoginCodeDongContent, v)
}

func (ct *cachedTmpl) StartupContent(v any) (*bytes.Buffer, error) {
	return ct.Render(tplStartupContent, v)
}
