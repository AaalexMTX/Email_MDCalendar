package checktime

import "time"

// Core 计时器内核
type Core struct {
	start     time.Time     // 开始时间
	frequency time.Duration // 发送频率
}

func NewCore() *Core {
	return &Core{}
}
