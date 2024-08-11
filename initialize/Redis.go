package initialize

import (
	"Glyphrz-go/global"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-redis/redis"
)

func InitRedis() {
	redisInfo := global.Settings.RedisInfo
	dsn := fmt.Sprintf("%s:%d", redisInfo.Host, redisInfo.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		color.Red("Redis连接错误")
		color.Yellow(err.Error())
	}
	global.Redis = client
}
