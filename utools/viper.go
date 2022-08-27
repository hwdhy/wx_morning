package utools

import (
	"github.com/spf13/viper"
	"log"
)

// ViperInit 初始化读取配置文件
func ViperInit() {
	viper.SetConfigName("config.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
}
