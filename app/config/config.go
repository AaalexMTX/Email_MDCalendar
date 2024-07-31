package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

// SetUpConfig 加载配置到内存
func (cfg *Core) SetUpConfig() error {
	// 初始化配置读取器
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("toml")
	// 指定配置文件路径(相对于main.go所在路径)
	vp.AddConfigPath("./app")
	// 查找并读取配置文件
	err := vp.ReadInConfig()
	if err != nil {
		var notFoundError viper.ConfigFileNotFoundError
		// 如果错误是因为找不到文件的报错
		if errors.As(err, &notFoundError) {
			fmt.Printf("file not exist\n")
		} else {
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
		return err
	}
	// 配置写入内存
	if err := vp.Unmarshal(&cfg); err != nil {
		fmt.Print(err.Error())
		return err
	}
	// 读取成功
	return nil
}
