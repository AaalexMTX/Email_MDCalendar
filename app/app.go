package app

import (
	"crypto/tls"
	"email_mdCalendar/app/config"
	"email_mdCalendar/app/message"
	"fmt"
	gomail "gopkg.in/mail.v2"
)

// Core 应用内核结构
type Core struct {
	cfg config.Core
	msg message.Core
	// logger
}

func Start() {
	// 初始化配置内核
	cfg := config.NewCore()
	// 配置读取
	if err := cfg.SetUpConfig(); err != nil {
		panic(err)
		return
	}

	// 设备消息内核
	msg := message.NewCore()
	msg.SetEmail()
	app := NewCore(cfg, msg)
	// 发送邮件
	if err := app.SendMessage(); err != nil {
		fmt.Printf("%s happen %s", "SendMessage", err.Error())
		return
	}
	fmt.Printf("send successful")
}

// NewCore 初始化应用内核
func NewCore(cfg *config.Core, msg *message.Core) *Core {
	return &Core{
		cfg: *cfg,
		msg: *msg,
	}
}

// SendMessage 发送邮件
func (c *Core) SendMessage() error {
	// 设置邮件内容
	msg := gomail.NewMessage()
	// 设置消息结构
	msg.SetHeader("From", c.cfg.Info.From)
	msg.SetHeader("To", c.cfg.Info.To)
	msg.SetHeader("Subject", c.msg.Subject)
	msg.SetBody("text/plain", c.msg.Body)
	// 设置连接信息
	d := gomail.NewDialer(c.cfg.Server.Smtp, c.cfg.Server.Port, c.cfg.Info.From, c.cfg.Info.AuthCode)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	// 建立smtp连接并发送消息
	if err := d.DialAndSend(msg); err != nil {
		fmt.Printf("%s happen %s", "DialAndSend", err.Error())
		return err
	}
	return nil
}
