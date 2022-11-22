package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	userProto "video/api/qvbilam/user/v1"
	"video/config"
)

var (
	DB               *gorm.DB
	ES               *elastic.Client
	RedisClient      *redis.Client
	ServerConfig     *config.ServerConfig
	UserServerClient userProto.UserClient
)
