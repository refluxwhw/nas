package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

var gLocalConfig = viper.New()

func LoadConfig(path string) error {
	// 使用 viper 加载配置文件
	gLocalConfig.SetConfigFile(path)
	if err := gLocalConfig.ReadInConfig(); err != nil {
		return fmt.Errorf("read config error: %v", err)
	}
	return nil
}

func GetConfig() *viper.Viper {
	return gLocalConfig
}
