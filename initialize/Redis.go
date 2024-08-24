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
	// 判断Redis是否在线，使用Ping Pong!
	_, err := client.Ping().Result()
	// 如果Redis不在线，直接Panic，Redis状态必须正常
	// Redis启动
	// redis-server 启动服务
	// redis-cli -h 127.0.0.1 -p 6379 连接Redis
	if err != nil {
		color.Red("Redis连接错误")
		panic(err)
		//color.Yellow(err.Error())
	}
	global.Redis = client
}
