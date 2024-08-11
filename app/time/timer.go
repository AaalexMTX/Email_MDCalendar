package checktime

import (
	"email_mdCalendar/app/config"
	"fmt"
	"time"
)

// StartTimer 开启定时器
func (c *Core) StartTimer(cfg *config.Core) {

}

// ConsistentTime 对齐第一次启动时间
func (c *Core) ConsistentTime(ch chan time.Time) bool {
	// 用当前时间与配置文件中的起始时间比较，获取初始计时器
	now := time.Now()
	// 如果当前时间大于配置文件中的起始时间, 认为首次发生时间已经过去，出错
	if now.After(c.start) {
		return false
	}
	// 首次校时定时器
	consistentTimer := time.NewTimer(c.start.Sub(now))
	// 将差时计时器的内容(差时计时器触发时间)，读出后再写入消息 channel中
	// 1. chan内容传递有点问题
	//ch <- <-consistentTimer.C

	// 2. ok
	//if sendTime, ok := <-consistentTimer.C; ok {
	//	ch <- sendTime
	//}
	//return true

	var ctime time.Time
	ctime = <-consistentTimer.C
	// 向通道发送校时后的时间，用作发送邮件的信号
	ch <- ctime
	fmt.Println("\nfirst consistent in --> ", ctime)
	return true
}

// CycleTime 循环间隔设定时间发送
func (c *Core) CycleTime() {

}

// LoadTimer 从配置文件的中解析出计时器的信息
func (c *Core) LoadTimer(cfg *config.Core) error {
	var err error

	// 解析配置文件中设置的起始时间
	/*
	 默认使用当前所在的时区解析时间！！！！
	 否则以 Asia/Shanghai 为例
	 由于默认使用 UTC，即GMT时间
	 解析出来的时间会较于实际时间滞后8小时
	*/
	c.start, err = time.ParseInLocation(time.DateTime, cfg.Timer.Starts, time.Now().Location())
	// 起始时间格式错误，解析失败
	if err != nil {
		return err
	}

	// TODO:解析配置文件中设置的发送间隔
	c.frequency = 12 * time.Second
	// 下次发送的日期

	return nil
}
