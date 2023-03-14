package sendto

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/vela-ssoc/backend-common/opurl"

	"github.com/vela-ssoc/backend-common/model"
)

// AlertConfigurer 运维中心告警平台配置
type AlertConfigurer interface {
	// AlertUnset 重置运维中心告警平台配置
	AlertUnset()

	// AlertConfig 获取运维中心告警平台配置
	AlertConfig() (*model.Alert, error)
}

// notifier 通知者
type notifier struct {
	Name          string `json:"name,omitempty"`           // Y 用户名
	Mobile        string `json:"mobile,omitempty"`         // N 用户手机号, 短信、微信需要(微信需要用户关注企业微信)
	Email         string `json:"email,omitempty"`          // N 用户邮箱、发邮件
	SerialNumber  string `json:"serial_number,omitempty"`  // N 工号、暂时可忽略
	NotifyMethods string `json:"notify_methods,omitempty"` // Y 告警方式，目前支持微信(weixin)、短信(sms)、邮件(email),可多选
}

// notifiers 通知者 slice
type notifiers []*notifier

// alertRequest 运维平台告警中心请求数据
// http://yunwei.eastmoney.com/docs/alert_api
type alertRequest struct {
	OriginName     string `json:"origin_name,omitempty"`     // Y 告警来源(发送告警的服务名称),需找服务人员事先添加
	AlertType      string `json:"alert_type,omitempty"`      // Y 告警类型,如 System
	AlertObject    string `json:"alert_object,omitempty"`    // Y 告警对象,如 cpu
	AlertAttribute string `json:"alert_attribute,omitempty"` // Y 告警字段,如 load
	Subject        string `json:"subject,omitempty"`         // Y 告警内容,(短信、微信内容；邮件标题)
	Notifier       string `json:"notifier,omitempty"`        // N 告警接收者,注意:字符串内容格式有要求,详细格式见notifier格式,
	Severity       string `json:"severity,omitempty"`        // N 告警级别,默认notice [disaster,high,middle,notice]
	Body           string `json:"body,omitempty"`            // N 邮件内容
	Project        string `json:"project,omitempty"`         // N 应用名称
	ProjectLevel   int    `json:"project_level,omitempty"`   // N 应用级别
	IP             string `json:"ip,omitempty"`              // N 告警对象ip地址，主要用于查询cmdb信息用于告警等
	URL            string `json:"url,omitempty"`             // N [文档未说明]
	AlertAt        string `json:"alert_at,omitempty"`        // N 告警时间,默认系统接收时间，format:2017-01-01 14:27:43
	Description    string `json:"description,omitempty"`     // N 告警备注
	Tags           string `json:"tags,omitempty"`            // N 告警标签，强大的功能支持,可定义形如(hostname=localhost,project=web)格式的任意多个内容
	notifiers      notifiers
}

// JSON 将数据转为 json 格式
func (ar alertRequest) JSON() (*bytes.Buffer, error) {
	if ar.Notifier == "" && len(ar.notifiers) != 0 {
		notice, err := json.Marshal(ar.notifiers)
		if err != nil {
			return nil, err
		}
		ar.Notifier = string(notice)
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(ar)

	return buf, err
}

// alertResponse 告警接口响应信息
type alertResponse struct {
	AlertID int64  `json:"alert_id"` // 告警接口调用成功会生成一个 ID
	Message string `json:"message"`  // 告警失败会返回一个错误原因
}

type AlertSender interface {
	EmailSender
	WechatSender
	SMSSender
	PhoneSender
}

func Alert(configure AlertConfigurer, client opurl.Client) AlertSender {
	return &alertClient{configure: configure, client: client}
}

type alertClient struct {
	configure AlertConfigurer
	client    opurl.Client
}

func (ac *alertClient) SendEmail(nos []string, title, content string) error {
	return ac.sendEmail(nos, title, content)
}

func (ac *alertClient) SendWechat(nos []string, title, content string) error {
	return ac.sendWechat(nos, title, content)
}

func (ac *alertClient) SendSMS(nos []string, content string) error {
	return ac.sendSMS(nos, content)
}

func (ac *alertClient) SendPhone(nos []string, content string) error {
	return ac.sendPhone(nos, content)
}

func (ac *alertClient) sendEmail(nos []string, title, content string) error {
	ntfs := make(notifiers, len(nos))
	for i, no := range nos {
		ntfs[i] = &notifier{Mobile: no, NotifyMethods: "email"}
	}

	req := &alertRequest{
		Subject:   title,
		Body:      content,
		notifiers: ntfs,
	}

	return ac.send(req)
}

func (ac *alertClient) sendWechat(nos []string, title, content string) error {
	ntfs := make(notifiers, len(nos))
	for i, no := range nos {
		ntfs[i] = &notifier{Mobile: no, NotifyMethods: "weixin"}
	}

	req := &alertRequest{
		Subject:   content,
		notifiers: ntfs,
	}

	return ac.send(req)
}

func (ac *alertClient) sendSMS(nos []string, content string) error {
	ntfs := make(notifiers, len(nos))
	for i, no := range nos {
		ntfs[i] = &notifier{Mobile: no, NotifyMethods: "sms"}
	}

	req := &alertRequest{
		Subject:   content,
		notifiers: ntfs,
	}

	return ac.send(req)
}

// sendPhone 发送电话通知
func (ac *alertClient) sendPhone(nos []string, content string) error {
	ntfs := make(notifiers, len(nos))
	for i, no := range nos {
		ntfs[i] = &notifier{Mobile: no, NotifyMethods: "call"}
	}

	req := &alertRequest{
		Subject:   content,
		notifiers: ntfs,
	}

	return ac.send(req)
}

// send 通过运维平台发送告警
func (ac *alertClient) send(r *alertRequest) error {
	cfg, err := ac.configure.AlertConfig()
	if err != nil {
		return err
	}
	r.OriginName = cfg.Origin
	r.AlertType = "ssoc"
	r.AlertObject = "ssoc"
	r.AlertAttribute = "ssoc"
	buf, err := r.JSON()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	header := make(http.Header, 4)
	if host := cfg.Host; host != "" {
		header.Set("Host", host)
	}
	res, err := ac.client.FetchRaw(ctx, http.MethodPost, cfg.Addr, nil, buf)
	if err != nil {
		return err
	}

	ret := new(alertResponse)
	err = json.NewDecoder(res.Body).Decode(ret)
	_ = res.Body.Close()
	if err == nil && ret.AlertID != 0 && ret.Message == "" {
		return nil
	}

	return errors.New(ret.Message)
}
