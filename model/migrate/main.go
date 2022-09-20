package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"video/global"
	"video/model"
)

func main() {
	db := DB()
	Migrate(db)
	Insert(db)
}

func InitElasticSearch() {
	host := "127.0.0.1"
	port := 9200

	url := elastic.SetURL(fmt.Sprintf("http://%s:%d", host, port))
	sniff := elastic.SetSniff(false)                             // 不将本地地址转换
	logger := log.New(os.Stdout, "elasticsearch", log.LstdFlags) // 设置日志输出位置
	var err error
	global.ES, err = elastic.NewClient(url, sniff, elastic.SetTraceLog(logger))
	if err != nil {
		zap.S().Panicf("连接es异常: %s", err.Error())
	}
}

func DB() *gorm.DB {
	InitElasticSearch()
	user := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	database := "qvbilam_video"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //不带表名
		},
	})
	if err != nil {
		panic(any(err))
	}
	return db
}

func Migrate(db *gorm.DB) {
	_ = db.AutoMigrate(
		&model.Region{},
		&model.Category{},
		&model.Style{},
		&model.StyleVideo{},
		&model.Video{},
		&model.Drama{},
		&model.DramaVideo{},
	)
}

func Insert(db *gorm.DB) {
	CreateCategory(db)
	CreateRegions(db)
	// todo 数据结构变动, 不适合插入
	//CreateEpisodes(db)
	//CreateVideo(db)
}
