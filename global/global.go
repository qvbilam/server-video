package global

import (
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"video/config"
)

var (
	DB           *gorm.DB
	ES           *elastic.Client
	ServerConfig config.ServerConfig
)
