package global

import (
	"Glyphrz-go/config"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Settings config.ServerConfig
	Log      *zap.Logger
	DB       *gorm.DB
	Redis    *redis.Client
)
