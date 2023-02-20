package evtrsk

import (
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/vela-ssoc/backend-common/model"
	"github.com/vela-ssoc/vela-manager/inward/plate"
	"github.com/vela-ssoc/vela-manager/outward/sendto"
	"gorm.io/gorm"
)

type Handler interface {
	Event(*model.Event) error
	Risk(*model.Risk) error
}

type Subscriber interface {
	Unset()
	Event(*model.Event) model.Recipients
	Risk(*model.Risk) model.Recipients
}

func NewHandle(
	db *gorm.DB,
	sub Subscriber,
	rend plate.Render,
	send sendto.Sender,
) Handler {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &notice{
		db:     db,
		sub:    sub,
		rend:   rend,
		send:   send,
		random: random,
	}
}

type notice struct {
	db     *gorm.DB      // 数据库
	sub    Subscriber    // 订阅者
	rend   plate.Render  // 模板渲染器
	send   sendto.Sender // 发送器
	random *rand.Rand    // 随机数生成器
}

func (ntc *notice) Event(evt *model.Event) error {
	if evt == nil {
		return nil
	}
	// 将产生的 event 保存到数据库
	if err := ntc.db.Create(evt).Error; err != nil || !evt.SendAlert {
		return err
	}

	evt.HaveRead = false
	secret := make([]byte, 16)
	ntc.random.Read(secret)
	evt.Secret = hex.EncodeToString(secret)

	// 获取 event 订阅者的通知方式
	var dong, email, wechat, sms, phone []string
	subs := ntc.sub.Event(evt)
	for _, sub := range subs {
		if sub.SendDong && sub.Dong != "" {
			dong = append(dong, sub.Dong)
		}
		if sub.SendEmail && sub.Email != "" {
			email = append(email, sub.Email)
		}
		if sub.SendWechat && sub.Wechat != "" {
			wechat = append(wechat, sub.Wechat)
		}
		if sub.SendSMS && sub.SMS != "" {
			sms = append(sms, sub.SMS)
		}
		if sub.SendPhone && sub.Phone != "" {
			phone = append(phone, sub.Phone)
		}
	}

	var ret SendError
	ret.Dong = ntc.dongEvent(evt, dong)
	ret.Email = ntc.emailEvent(evt, email)
	ret.Wechat = ntc.wechatEvent(evt, wechat)
	ret.SMS = ntc.smsEvent(evt, sms)
	ret.Phone = ntc.phoneEvent(evt, phone)

	if ret.Dong == nil &&
		ret.Email == nil &&
		ret.Wechat == nil &&
		ret.SMS == nil &&
		ret.Phone == nil {
		return nil
	}

	return ret
}

func (ntc *notice) Risk(rsk *model.Risk) error {
	if rsk == nil {
		return nil
	}
	// 将产生的 event 保存到数据库
	if err := ntc.db.Create(rsk).Error; err != nil || !rsk.SendAlert {
		return err
	}

	// 获取 event 订阅者的通知方式
	var dong, email, wechat, sms, phone []string
	subs := ntc.sub.Risk(rsk)
	for _, sub := range subs {
		if sub.SendDong && sub.Dong != "" {
			dong = append(dong, sub.Dong)
		}
		if sub.SendEmail && sub.Email != "" {
			email = append(email, sub.Email)
		}
		if sub.SendWechat && sub.Wechat != "" {
			wechat = append(wechat, sub.Wechat)
		}
		if sub.SendSMS && sub.SMS != "" {
			sms = append(sms, sub.SMS)
		}
		if sub.SendPhone && sub.Phone != "" {
			phone = append(phone, sub.Phone)
		}
	}

	var ret SendError
	ret.Dong = ntc.dongRisk(rsk, dong)
	ret.Email = ntc.emailRisk(rsk, email)
	ret.Wechat = ntc.wechatRisk(rsk, wechat)
	ret.SMS = ntc.smsRisk(rsk, sms)
	ret.Phone = ntc.phoneRisk(rsk, phone)

	if ret.Dong == nil &&
		ret.Email == nil &&
		ret.Wechat == nil &&
		ret.SMS == nil &&
		ret.Phone == nil {
		return nil
	}

	return ret
}

func (ntc *notice) dongEvent(evt *model.Event, tos []string) error {
	if len(tos) == 0 {
		return nil
	}
	tit, err := ntc.rend.EventDongTitle(evt)
	if err != nil {
		return err
	}
	ctc, err := ntc.rend.EventDongContent(evt)
	if err != nil {
		return err
	}

	title := tit.String()
	content := ctc.String()

	return ntc.send.SendDong(tos, title, content)
}

func (ntc *notice) emailEvent(evt *model.Event, tos []string) error {
	if len(tos) == 0 {
		return nil
	}
	tit, err := ntc.rend.EventEmailTitle(evt)
	if err != nil {
		return err
	}
	ctc, err := ntc.rend.EventEmailContent(evt)
	if err != nil {
		return err
	}

	title := tit.String()
	content := ctc.String()

	return ntc.send.SendEmail(tos, title, content)
}

func (ntc *notice) wechatEvent(evt *model.Event, tos []string) error {
	if len(tos) == 0 {
		return nil
	}
	tit, err := ntc.rend.EventWechatTitle(evt)
	if err != nil {
		return err
	}
	ctc, err := ntc.rend.EventWechatContent(evt)
	if err != nil {
		return err
	}

	title := tit.String()
	content := ctc.String()

	return ntc.send.SendWechat(tos, title, content)
}

func (ntc *notice) smsEvent(evt *model.Event, tos []string) error {
	if len(tos) == 0 {
		return nil
	}
	ctc, err := ntc.rend.EventSMSContent(evt)
	if err != nil {
		return err
	}
	content := ctc.String()

	return ntc.send.SendSMS(tos, content)
}

func (ntc *notice) phoneEvent(evt *model.Event, tos []string) error {
	if len(tos) == 0 {
		return nil
	}
	ctc, err := ntc.rend.EventPhoneContent(evt)
	if err != nil {
		return err
	}
	content := ctc.String()

	return ntc.send.SendPhone(tos, content)
}

func (ntc *notice) dongRisk(rsk *model.Risk, tos []string) error {
	if len(tos) == 0 {
		return nil
	}
	tit, err := ntc.rend.RiskDongTitle(rsk)
	if err != nil {
		return err
	}
	ctc, err := ntc.rend.RiskDongContent(rsk)
	if err != nil {
		return err
	}

	title := tit.String()
	content := ctc.String()

	return ntc.send.SendDong(tos, title, content)
}

func (ntc *notice) emailRisk(rsk *model.Risk, tos []string) error {
	if len(tos) == 0 {
		return nil
	}
	tit, err := ntc.rend.RiskEmailTitle(rsk)
	if err != nil {
		return err
	}
	ctc, err := ntc.rend.RiskEmailContent(rsk)
	if err != nil {
		return err
	}

	title := tit.String()
	content := ctc.String()

	return ntc.send.SendEmail(tos, title, content)
}

func (ntc *notice) wechatRisk(rsk *model.Risk, tos []string) error {
	if len(tos) == 0 {
		return nil
	}
	tit, err := ntc.rend.RiskWechatTitle(rsk)
	if err != nil {
		return err
	}
	ctc, err := ntc.rend.RiskWechatContent(rsk)
	if err != nil {
		return err
	}

	title := tit.String()
	content := ctc.String()

	return ntc.send.SendWechat(tos, title, content)
}

func (ntc *notice) smsRisk(rsk *model.Risk, tos []string) error {
	if len(tos) == 0 {
		return nil
	}
	ctc, err := ntc.rend.RiskSMSContent(rsk)
	if err != nil {
		return err
	}
	content := ctc.String()

	return ntc.send.SendSMS(tos, content)
}

func (ntc *notice) phoneRisk(rsk *model.Risk, tos []string) error {
	if len(tos) == 0 {
		return nil
	}
	ctc, err := ntc.rend.RiskPhoneContent(rsk)
	if err != nil {
		return err
	}
	content := ctc.String()

	return ntc.send.SendPhone(tos, content)
}
