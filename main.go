package main

import (
	"email_mdCalendar/app"
	"fmt"
)

// main
// 日志文件读取
// 兼容 md生成程序（用go重写）
// 定时发送系统
func main() {
	// func test
	fmt.Print(appName)
	// 启动程序
	if err := app.Start(); err != nil {
		fmt.Printf("程序启动失败: %v\n", err.Error())
	}
	fmt.Print("Quit successful\n")
}

var appName string = `
	     ████████╗██╗███╗   ███╗███████╗██████╗         ███████╗███╗   ███╗ █████╗ ██╗██╗     
	    ╚══██╔══╝██║████╗ ████║██╔════╝██╔══██╗        ██╔════╝████╗ ████║██╔══██╗██║██║     
	      ██║   ██║██╔████╔██║█████╗  ██║  ██║        █████╗  ██╔████╔██║███████║██║██║     
	     ██║   ██║██║╚██╔╝██║██╔══╝  ██║  ██║        ██╔══╝  ██║╚██╔╝██║██╔══██║██║██║     
	    ██║   ██║██║ ╚═╝ ██║███████╗██████╔╝███████╗███████╗██║ ╚═╝ ██║██║  ██║██║███████╗
	   ╚═╝   ╚═╝╚═╝     ╚═╝╚══════╝╚═════╝ ╚══════╝╚══════╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝╚══════╝
`
