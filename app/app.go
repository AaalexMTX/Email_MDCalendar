package app

import (
	"crypto/tls"
	"email_mdCalendar/app/config"
	"email_mdCalendar/app/message"
	checktime "email_mdCalendar/app/time"
	"fmt"
	gomail "gopkg.in/mail.v2"
)

// Core 应用内核结构
type Core struct {
	cfg        *config.Core
	msg        *message.Core
	timer      *checktime.Core
	msgChan    <-chan struct{}
	emailCount int
	// logger
}

// Start 启动应用程序
func Start() error {
	// 创建应用内核
	app := NewCore()
	// 读取配置文件并初始化内核设置
	if err := app.SetUpConfig(); err != nil {
		return err
	}
	// TODO: 放进 启动邮件发送定时器
	app.timer.StartTimer(app.cfg)

	//校验时间，对齐第一次发送时间
	if app.timer.ConsistentTime() {
		// 发送邮件
		if err := app.SendMessage(); err != nil {
			return err
		}
		app.emailCount++
	} else {
		fmt.Println("time Out")
	}
	// 发送次数未到 或 需一直发送，则循环
	for ; app.emailCount <= app.cfg.Timer.Count || app.cfg.Timer.Count < 0; app.emailCount++ {
		// TODO: 开启循环计时器，等待发送邮件
		app.timer.CycleTime()
	}

	// 正确结束
	return nil
}

// NewCore 初始化应用内核
func NewCore() *Core {
	// 初始化配置内核
	cfg := config.NewCore()
	// 初始化设备消息内核
	msg := message.NewCore()
	// 初始化时间内核
	t := checktime.NewCore()
	// 返回应用内核
	return &Core{
		msgChan:    make(chan struct{}),
		emailCount: 0,
		cfg:        cfg,
		msg:        msg,
		timer:      t,
	}
}

// SetUpConfig 读入设置内核结构
func (c *Core) SetUpConfig() error {
	// 配置文件读取
	if err := c.cfg.SetUpConfig(); err != nil {
		return err
	}

	// 设置消息体
	c.msg.SetEmail()

	// 读取解析定时器
	if err := c.timer.LoadTimer(c.cfg); err != nil {
		return err
	}

	return nil
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
		return err
	}
	fmt.Printf("Email -- %02d -- send successful\n", c.emailCount)
	return nil
}
