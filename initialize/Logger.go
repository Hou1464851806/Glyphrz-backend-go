package initialize

import (
	"Glyphrz-go/global"
	"Glyphrz-go/utils"
	"fmt"
	"go.uber.org/zap"
)

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{
		fmt.Sprintf("%slog_%s", global.Settings.LogAddress, utils.GetTodayFormatTime()),
		"stdout",
	}
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
	global.Log = logger
}
