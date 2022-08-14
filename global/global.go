package global

import (
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	userProto "video/api/user/pb"
	"video/config"
)

var (
	DB               *gorm.DB
	ES               *elastic.Client
	ServerConfig     config.ServerConfig
	UserServerClient userProto.UserClient
)
