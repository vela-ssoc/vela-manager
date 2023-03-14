package notification

import "fmt"

// dongCodes 咚咚平台定义的错误码
var dongCodes = map[string]string{
	"4000009": "未登陆",
	"4000001": "参数无效",
	"4000002": "权限不足",
	"4000003": "验证码错误",
	"4000004": "登陆频率过高",
	"4000005": "验证码已失效",
	"4000006": "没有找到资源",
	"4000007": "系统错误，请联系管理员",
	"4000008": "无权访问",
	"4000010": "上传提交的数据已经在审批中",
	"4000012": "消息类型不匹配",
	"4000013": "群必须包含 3 个以上成员",
	"4000014": "账号已失效",
	"4000015": "认证失败",
	"5000001": "用户不存在",
	"5000002": "用户名密码不匹配",
	"5000003": "没有权限登录",
	"5000004": "无该用户实名信息",
	"5000005": "文件格式不正确",
	"5000006": "用户越权操作",
	"5000009": "服务访问频率受限",
	"5000010": "服务 IP 不在白名单内",
}

// DongError 发送咚咚错误
type DongError struct {
	Code  string // 咚咚服务返回的业务状态码
	Msg   string // 消息
	Cause string // 原因
}

// Error error
func (de *DongError) Error() string {
	return fmt.Sprintf("咚咚服务器返回错误 [%s]: %s, %s", de.Code, de.Msg, de.Cause)
}

// dongRequest 咚咚通知请求报文
type dongRequest struct {
	UserIDs string `json:"userIds"` // 消息接收用户，多个 以,隔开（groupIds 为 空时必填）
	Title   string `json:"title"`   // 标题（必填）长度：300 个字符（150 个中文）
	Detail  string `json:"detail"`  // 卡片消息详细（必填）长度限制：2000 个字符（中文：1000 个）支持 html 标签
}

// dongResponse 咚咚通知请求报文
type dongResponse struct {
	Code string `json:"code"` // 请求返回码
	Msg  string `json:"msg"`  // 请求返回消息
}

// Error 判断响应是否包含错误
func (dr dongResponse) Error() error {
	code := dr.Code
	if code == "200" {
		return nil
	}
	cause := dongCodes[code]

	return &DongError{
		Code:  code,
		Msg:   dr.Msg,
		Cause: cause,
	}
}
