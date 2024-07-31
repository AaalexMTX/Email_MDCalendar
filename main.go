package main

import (
	"crypto/tls"
	"email_mdCalendar/config"
	"fmt"
	gomail "gopkg.in/mail.v2"
)

// main
// 日志文件读取
// 兼容 md生成程序（用go重写）
// 定时发送系统
func main() {
	// 配置读取
	cfg := config.CfgInput{}
	err := config.SetUpConfig(&cfg)
	if err != nil {
		panic(err)
	}

	msg := gomail.NewMessage()
	// 设置消息结构
	msg.SetHeader("From", cfg.Info.From)
	msg.SetHeader("To", cfg.Info.To)
	msg.SetHeader("Subject", "contents-2")
	msg.SetBody("text/plain", "test2--markdown")
	// 设置连接信息
	d := gomail.NewDialer(cfg.Server.Smtp, cfg.Server.Port, cfg.Info.From, cfg.Info.AuthCode)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	// 建立smtp连接并发送消息
	if err := d.DialAndSend(msg); err != nil {
		fmt.Print(err)
		panic(err)
	}
	fmt.Printf("send successful")
}
