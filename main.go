package main

import (
	"Glyphrz-go/global"
	"Glyphrz-go/initialize"
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"time"
)

func main() {
	// 初始化配置
	initialize.InitConfig()
	// 初始化路由
	r := initialize.Router()
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitMysqlDB()
	// 初始化Redis
	initialize.InitRedis()
	// 测试redis
	global.Redis.Set("test", "testValue", time.Second)
	//time.Sleep(time.Second * 2)
	value := global.Redis.Get("test")
	color.Blue(value.Val())
	// 启动服务
	if err := r.Run(fmt.Sprintf(":%d", global.Settings.Port)); err != nil {
		zap.L().Info("Server setup function", zap.String("error", "Start error!"))
	}
}
