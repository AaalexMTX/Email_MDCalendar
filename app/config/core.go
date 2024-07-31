package config

// Core  配置文件的映射结构
type Core struct {
	Server server `mapstructure:"server"`
	Info   info   `mapstructure:"info"`
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

// NewCore 初始化配置结构
func NewCore() *Core {
	return &Core{}
}
