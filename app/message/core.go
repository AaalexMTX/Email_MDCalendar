package message

import "time"

// Core 邮件内容信息内核
type Core struct {
	Subject string    // 邮件主题
	Body    string    // 邮件内容
	Time    time.Time // 发送时间
}

// NewCore 创建邮件信息内核
func NewCore() *Core {
	return &Core{}
}
