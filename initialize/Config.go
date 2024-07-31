package initialize

import (
	"Glyphrz-go/config"
	"Glyphrz-go/global"
	"github.com/spf13/viper"
)

func InitConfig() {
	//实例化viper
	v := viper.New()
	//设置Config文件目录
	v.SetConfigFile("./Config.yaml")
	//读取Config文件
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 创建serverConfig结构体，用于实例化
	serverConfig := config.ServerConfig{}
	// 反序列化Config
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	// 存储在全局变量中
	global.Settings = serverConfig
}
