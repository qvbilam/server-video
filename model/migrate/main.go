package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"video/model"
)

func main() {
	db := DB()
	Migrate(db)
	Insert(db)
}

func DB() *gorm.DB {
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
		&model.Episodes{},
		&model.Barrage{},
	)
}

func Insert(db *gorm.DB) {
	CreateCategory(db)
	CreateEpisodes(db)
	CreateRegions(db)
	CreateVideo(db)
}
