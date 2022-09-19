package global

import (
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	userProto "video/api/qvbilam/user/v1"
	"video/config"
)

var (
	DB               *gorm.DB
	ES               *elastic.Client
	ServerConfig     config.ServerConfig
	UserServerClient userProto.UserClient
)
