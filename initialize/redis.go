package initialize

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"video/global"
)

func InitRedis() {
	addr := fmt.Sprintf("%s:%d", global.ServerConfig.RedisConfig.Host, global.ServerConfig.RedisConfig.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: global.ServerConfig.RedisConfig.Password,
		DB:       global.ServerConfig.RedisConfig.Database,
	})

	global.RedisClient = rdb
}
