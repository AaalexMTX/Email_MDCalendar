package checktime

import "time"

// Core 计时器内核
type Core struct {
	start     time.Time     // 开始时间
	frequency time.Duration // 发送频率
	calendar  Calendar      // 日历
}

// Calendar 日历计算结构
type Calendar struct {
	Date Date
}

// Date 日期
type Date struct {
	Year  int
	Month int
}

func NewCore() *Core {
	return &Core{}
}
