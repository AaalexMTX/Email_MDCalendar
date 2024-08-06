package config

// Core  配置文件的映射结构
type Core struct {
	Info   info   `mapstructure:"info"`
	Server server `mapstructure:"server"`
	Timer  timer  `mapstructure:"timer"`
}

// server 邮件服务器配置
type server struct {
	Smtp string `mapstructure:"smtp"`
	Port int    `mapstructure:"port"`
}

// info 邮件发送配置
type info struct {
	From     string `mapstructure:"from"`      // 发件人
	To       string `mapstructure:"to"`        // 接收人
	AuthCode string `mapstructure:"auth_code"` // 邮箱授权码
}

// timer 定时邮件任务配置
type timer struct {
	Enable    bool     `mapstructure:"timed_task"`    // 是否启用定时任务
	Starts    string   `mapstructure:"start_time"`    // 定时任务开始时间
	Frequency interval `mapstructure:"time_interval"` // 定时任务频率
}

type interval struct {
	Year   string `mapstructure:"year"`
	Month  string `mapstructure:"month"`
	Day    string `mapstructure:"day"`
	Hour   string `mapstructure:"hour"`
	Minute string `mapstructure:"minute"`
	Second string `mapstructure:"second"`
}

// NewCore 初始化配置结构
func NewCore() *Core {
	return &Core{}
}
