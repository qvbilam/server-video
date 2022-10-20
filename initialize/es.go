package initialize

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"log"
	"os"
	"video/global"
	"video/model/doc"
)

func InitElasticSearch() {
	host := global.ServerConfig.ESConfig.Host
	port := global.ServerConfig.ESConfig.Port

	url := elastic.SetURL(fmt.Sprintf("http://%s:%d", host, port))
	sniff := elastic.SetSniff(false) // 不将本地地址转换
	var err error
	logger := log.New(os.Stdout, "elasticsearch", log.LstdFlags) // 设置日志输出位置
	global.ES, err = elastic.NewClient(url, sniff, elastic.SetTraceLog(logger))
	//global.ES, err = elastic.NewClient(url, sniff)
	if err != nil {
		zap.S().Panicf("连接es异常: %s", err.Error())
	}

	// 创建 video mapping
	createVideoIndex()
	createDramaIndex()
}

func createVideoIndex() {
	exists, err := global.ES.IndexExists(doc.Video{}.GetIndexName()).Do(context.Background())
	if err != nil {
		zap.S().Panicf("视频索引异常: %s", err)
	}
	if !exists { // 创建索引
		createIndex, err := global.ES.
			CreateIndex(doc.Video{}.GetIndexName()).
			BodyString(doc.Video{}.GetMapping()).
			Do(context.Background())
		if err != nil {
			zap.S().Panicf("创建视频索引异常: %s", err)
		}

		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
}

func createDramaIndex() {
	exists, err := global.ES.IndexExists(doc.Drama{}.GetIndexName()).Do(context.Background())
	if err != nil {
		zap.S().Panicf("视频索引异常: %s", err)
	}
	if !exists { // 创建索引
		createIndex, err := global.ES.
			CreateIndex(doc.Drama{}.GetIndexName()).
			BodyString(doc.Drama{}.GetMapping()).
			Do(context.Background())
		if err != nil {
			zap.S().Panicf("创建剧集索引异常: %s", err)
		}

		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
}
