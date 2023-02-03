package sendto

import "errors"

var (
	ErrUnimplementedDong   = errors.New("尚未实现咚咚通知功能")
	ErrUnimplementedEmail  = errors.New("尚未实现邮件通知功能")
	ErrUnimplementedWechat = errors.New("尚未实现微信通知功能")
	ErrUnimplementedSMS    = errors.New("尚未实现短信通知功能")
	ErrUnimplementedPhone  = errors.New("尚未实现电话通知功能")
)

type DongSender interface {
	// SendDong 向咚咚发送通知
	SendDong(dongs []string, title, content string) error
}

type EmailSender interface {
	// SendEmail 发送邮件通知
	SendEmail(tos []string, title, content string) error
}

type WechatSender interface {
	// SendWechat 发送微信通知
	SendWechat(nos []string, title, content string) error
}

type SMSSender interface {
	// SendSMS 发送短信通知
	SendSMS(nos []string, content string) error
}

type PhoneSender interface {
	// SendPhone 发送电话通知
	SendPhone(nos []string, content string) error
}

// Sender 发送通知
type Sender interface {
	DongSender
	EmailSender
	WechatSender
	SMSSender
	PhoneSender
}

type PostmanOption func(*postman)

func WithDong(dong DongSender) PostmanOption {
	return func(ptm *postman) {
		ptm.dong = dong
	}
}

func WithEmail(email EmailSender) PostmanOption {
	return func(ptm *postman) {
		ptm.email = email
	}
}

func WithWechat(wechat WechatSender) PostmanOption {
	return func(ptm *postman) {
		ptm.wechat = wechat
	}
}

func WithSMS(sms SMSSender) PostmanOption {
	return func(ptm *postman) {
		ptm.sms = sms
	}
}

func WithPhone(phone PhoneSender) PostmanOption {
	return func(ptm *postman) {
		ptm.phone = phone
	}
}

func Postman(opts ...PostmanOption) Sender {
	ptm := new(postman)
	for _, opt := range opts {
		opt(ptm)
	}

	return ptm
}

type postman struct {
	dong   DongSender
	email  EmailSender
	wechat WechatSender
	sms    SMSSender
	phone  PhoneSender
}

func (ptm *postman) SendDong(dongs []string, title, content string) error {
	if dong := ptm.dong; dong != nil {
		return dong.SendDong(dongs, title, content)
	}
	return ErrUnimplementedDong
}

func (ptm *postman) SendEmail(tos []string, title, content string) error {
	if email := ptm.email; email != nil {
		return email.SendEmail(tos, title, content)
	}
	return ErrUnimplementedEmail
}

func (ptm *postman) SendWechat(nos []string, title, content string) error {
	if wechat := ptm.wechat; wechat != nil {
		return wechat.SendWechat(nos, title, content)
	}
	return ErrUnimplementedWechat
}

func (ptm *postman) SendSMS(nos []string, content string) error {
	if sms := ptm.sms; sms != nil {
		return sms.SendSMS(nos, content)
	}
	return ErrUnimplementedSMS
}

func (ptm *postman) SendPhone(nos []string, content string) error {
	if phone := ptm.phone; phone != nil {
		return phone.SendPhone(nos, content)
	}
	return ErrUnimplementedPhone
}
