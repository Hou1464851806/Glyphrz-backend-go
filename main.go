package main

import (
	"Glyphrz-go/global"
	"Glyphrz-go/initialize"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	// 初始化配置
	initialize.InitConfig()
	// 初始化路由
	r := initialize.Router()
	// 初始化日志
	initialize.InitLogger()
	// 启动服务
	if err := r.Run(fmt.Sprintf(":%d", global.Settings.Port)); err != nil {
		zap.L().Info("Server setup function", zap.String("error", "Start error!"))
	}
}
