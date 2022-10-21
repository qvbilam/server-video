package api

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"testing"
	proto "video/api/qvbilam/video/v1"
)

var Client proto.VideoClient
var DramaClient proto.DramaClient
var ESClient *elastic.Client
var DBClient *gorm.DB

func initElasticSearch() {
	host := "127.0.0.1"
	port := 9200

	url := elastic.SetURL(fmt.Sprintf("http://%s:%d", host, port))
	sniff := elastic.SetSniff(false)                             // 不将本地地址转换
	logger := log.New(os.Stdout, "elasticsearch", log.LstdFlags) // 设置日志输出位置
	var err error
	ESClient, err = elastic.NewClient(url, sniff, elastic.SetTraceLog(logger))
	if err != nil {
		zap.S().Panicf("连接es异常: %s", err.Error())
	}
}

func initDBClient() {
	initElasticSearch()
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
	DBClient = db
}

func initClient() {
	port := 9802
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", port), grpc.WithInsecure())
	if err != nil {
		panic(any(err))
	}

	Client = proto.NewVideoClient(conn)
	DramaClient = proto.NewDramaClient(conn)
}

// 创建视频
func TestVideoServer_Create(t *testing.T) {
	initClient()

	response, err := Client.Create(context.Background(), &proto.UpdateVideoRequest{
		UserId:         1,
		CategoryId:     1,
		Name:           "测试视频",
		Introduce:      "测试视频简介",
		Icon:           "测试视频图标",
		HorizontalIcon: "测试视频纵版图标",
		DramaId:        3,
		Episode:        1,
	})

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(response)
}

// 获取视频
func TestVideoServer_Get(t *testing.T) {
	initClient()
	response, err := Client.Get(context.Background(), &proto.SearchVideoRequest{
		Keyword: "测试",
		Page:    1,
		PerPage: 10,
	})
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(response)
}

func TestVideoServer_Update(t *testing.T) {

	sorts := []string{
		"-play_count",
		"play_count",
	}

	for _, s := range sorts {
		if string(s[0]) == "-" {
			fmt.Printf("sort type: %s;", string(s[0]))
			fmt.Printf("sort field: %s\n", s[0:])
		} else {
			fmt.Printf("sort type: %s;", string(s[0]))
			fmt.Printf("sort field: %s\n", s[0:])
		}
	}

}
