package message

import (
	"time"
)

// SetEmail 设置邮件信息
func (c *Core) SetEmail() {
	_ = c.SetSubject()
	_ = c.SetBody()
	_ = c.SetTime()
}

// SetTime setTime 设置邮件发送时间
func (c *Core) SetTime() error {
	c.Time = time.Now()
	return nil
}

// SetSubject setSubject 设置邮件主题
func (c *Core) SetSubject() error {
	c.Subject = "2024-08 Markdown Calendar"
	return nil
}

// SetBody setBody 设置邮件内容
func (c *Core) SetBody() error {
	data := "normal text test"
	c.Body = data
	return nil
}
