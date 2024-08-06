package app

import (
	"crypto/tls"
	"email_mdCalendar/app/config"
	"email_mdCalendar/app/message"
	"email_mdCalendar/app/time"
	"fmt"
	gomail "gopkg.in/mail.v2"
)

// Core 应用内核结构
type Core struct {
	msgChan <-chan struct{}
	cfg     *config.Core
	msg     *message.Core
	timer   *time.Core
	// logger
}

// Start 启动应用程序
func Start() error {
	app := NewCore()
	// 配置读取
	if err := app.cfg.SetUpConfig(); err != nil {
		panic(err)
		return err
	}

	// 设置消息体
	app.msg.SetEmail()
	// 发送邮件
	if err := app.SendMessage(); err != nil {
		return err
	}
	fmt.Printf("send successful")
	return nil
}

// NewCore 初始化应用内核
func NewCore() *Core {
	// 初始化配置内核
	cfg := config.NewCore()
	// 初始化设备消息内核
	msg := message.NewCore()
	// 初始化时间内核
	t := time.NewCore()
	return &Core{
		msgChan: make(chan struct{}),
		cfg:     cfg,
		msg:     msg,
		timer:   t,
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
