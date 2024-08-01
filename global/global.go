package global

import (
	"Glyphrz-go/config"
	"go.uber.org/zap"
)

var (
	Settings config.ServerConfig
	Log      *zap.Logger
)
