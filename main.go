package main

import (
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/mail.v2"
)

type configInput struct {
	from     string
	to       string
	authCode string
}

const (
	from               = "2323426610@qq.com"
	to                 = "2323426610@qq.com"
	authorization_code = ""
)

func main() {
	msg := gomail.NewMessage()
	// 设置消息结构
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "contents")
	msg.SetBody("text/plain", "test")
	// 设置连接信息
	d := gomail.NewDialer("smtp.qq.com", 465, from, authorization_code)
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
