package sendto

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/vela-ssoc/manager/libkit/httpclient"
	"github.com/vela-ssoc/manager/model"
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
	Name          string `json:"name"`           // Y 用户名
	Mobile        string `json:"mobile"`         // N 用户手机号, 短信、微信需要(微信需要用户关注企业微信)
	Email         string `json:"email"`          // N 用户邮箱、发邮件
	SerialNumber  string `json:"serial_number"`  // N 工号、暂时可忽略
	NotifyMethods string `json:"notify_methods"` // Y 告警方式，目前支持微信(weixin)、短信(sms)、邮件(email),可多选
}

// notifiers 通知者 slice
type notifiers []*notifier

// alertRequest 运维平台告警中心请求数据
// http://yunwei.eastmoney.com/docs/alert_api
type alertRequest struct {
	OriginName     string `json:"origin_name"`     // Y 告警来源(发送告警的服务名称),需找服务人员事先添加
	AlertType      string `json:"alert_type"`      // Y 告警类型,如System
	AlertObject    string `json:"alert_object"`    // Y 告警对象,如cpu
	AlertAttribute string `json:"alert_attribute"` // Y 告警字段,如 load
	Subject        string `json:"subject"`         // Y 告警内容,(短信、微信内容；邮件标题)
	Notifier       string `json:"notifier"`        // N 告警接收者,注意:字符串内容格式有要求,详细格式见notifier格式,
	Severity       string `json:"severity"`        // N 告警级别,默认notice [disaster,high,middle,notice]
	Body           string `json:"body"`            // N 邮件内容
	Project        string `json:"project"`         // N 应用名称
	ProjectLevel   int    `json:"project_level"`   // N 应用级别
	IP             string `json:"ip"`              // N 告警对象ip地址，主要用于查询cmdb信息用于告警等
	URL            string `json:"url"`             // N [文档未说明]
	AlertAt        string `json:"alert_at"`        // N 告警时间,默认系统接收时间，format:2017-01-01 14:27:43
	Description    string `json:"description"`     // N 告警备注
	Tags           string `json:"tags"`            // N 告警标签，强大的功能支持,可定义形如(hostname=localhost,project=web)格式的任意多个内容
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

func Alert() {
}

type alertClient struct {
	configure AlertConfigurer
	client    httpclient.Client
}

func (ac *alertClient) sendEmail(nos []string, title, content string) error {
	ntfs := make(notifiers, len(nos))
	for _, no := range nos {
		ntfs = append(ntfs, &notifier{Email: no, NotifyMethods: "email"})
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
	for _, no := range nos {
		ntfs = append(ntfs, &notifier{Mobile: no, NotifyMethods: "weixin"})
	}

	req := &alertRequest{
		Subject:   content,
		notifiers: ntfs,
	}

	return ac.send(req)
}

func (ac *alertClient) sendSMS(nos []string, content string) error {
	ntfs := make(notifiers, len(nos))
	for _, no := range nos {
		ntfs = append(ntfs, &notifier{Mobile: no, NotifyMethods: "sms"})
	}

	req := &alertRequest{
		Subject:   content,
		notifiers: ntfs,
	}

	return ac.send(req)
}

func (ac *alertClient) sendPhone(nos []string, content string) error {
	ntfs := make(notifiers, len(nos))
	for _, no := range nos {
		ntfs = append(ntfs, &notifier{Mobile: no, NotifyMethods: "call"})
	}

	req := &alertRequest{
		Subject:   content,
		notifiers: ntfs,
	}

	return ac.send(req)
}

func (ac *alertClient) send(r *alertRequest) error {
	cfg, err := ac.configure.AlertConfig()
	if err != nil {
		return err
	}
	r.OriginName = cfg.Origin
	buf, err := r.JSON()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, cfg.Addr, buf)
	if host := cfg.Host; host != "" {
		req.Host = host
	}

	res, err := ac.client.Fetch(req)
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
