package notification

import (
	"bytes"
	"encoding/json"
	"strings"
)

// alertRequest 运维平台告警中心请求数据
// http://yunwei.eastmoney.com/docs/alert_api
type alertRequest struct {
	OriginName     string    `json:"origin_name,omitempty"`     // Y 告警来源(发送告警的服务名称),需找服务人员事先添加
	AlertType      string    `json:"alert_type,omitempty"`      // Y 告警类型,如 System
	AlertObject    string    `json:"alert_object,omitempty"`    // Y 告警对象,如 cpu
	AlertAttribute string    `json:"alert_attribute,omitempty"` // Y 告警字段,如 load
	Subject        string    `json:"subject,omitempty"`         // Y 告警内容,(短信、微信内容；邮件标题)
	Notifier       notifiers `json:"notifier,omitempty"`        // N 告警接收者,注意:字符串内容格式有要求,详细格式见notifier格式,
	Severity       string    `json:"severity,omitempty"`        // N 告警级别,默认notice [disaster,high,middle,notice]
	Body           string    `json:"body,omitempty"`            // N 邮件内容
	Project        string    `json:"project,omitempty"`         // N 应用名称
	ProjectLevel   int       `json:"project_level,omitempty"`   // N 应用级别
	IP             string    `json:"ip,omitempty"`              // N 告警对象ip地址，主要用于查询cmdb信息用于告警等
	URL            string    `json:"url,omitempty"`             // N [文档未说明]
	AlertAt        string    `json:"alert_at,omitempty"`        // N 告警时间,默认系统接收时间，format:2017-01-01 14:27:43
	Description    string    `json:"description,omitempty"`     // N 告警备注
	Tags           string    `json:"tags,omitempty"`            // N 告警标签，强大的功能支持,可定义形如(hostname=localhost,project=web)格式的任意多个内容
}

// notifier 通知者
type notifier struct {
	Name          string       `json:"name,omitempty"`           // Y 用户名
	Mobile        string       `json:"mobile,omitempty"`         // N 用户手机号, 短信、微信需要(微信需要用户关注企业微信)
	Email         string       `json:"email,omitempty"`          // N 用户邮箱、发邮件
	SerialNumber  string       `json:"serial_number,omitempty"`  // N 工号、暂时可忽略
	NotifyMethods stringsComma `json:"notify_methods,omitempty"` // Y 告警方式，目前支持微信(weixin)、短信(sms)、邮件(email),可多选
}

type stringsComma []string

func (sc stringsComma) MarshalText() ([]byte, error) {
	dat := strings.Join(sc, ",")
	return []byte(dat), nil
}

// notifiers 通知者 slice
type notifiers []*notifier

func (ns notifiers) MarshalText() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.WriteByte('[')
	max := len(ns) - 1
	for i, n := range ns {
		dat, err := json.Marshal(n)
		if err != nil {
			return nil, err
		}
		buf.Write(dat)
		if i < max {
			buf.WriteByte(',')
		}
	}

	buf.WriteByte(']')

	return buf.Bytes(), nil
}
