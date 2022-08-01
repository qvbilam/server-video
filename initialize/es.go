package initialize

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"log"
	"os"
	"video/global"
)

func InitElasticSearch() {
	host := global.ServerConfig.ESConfig.Host
	port := global.ServerConfig.ESConfig.Port

	url := elastic.SetURL(fmt.Sprintf("http://%s:%d", host, port))
	sniff := elastic.SetSniff(false)                             // 不将本地地址转换
	logger := log.New(os.Stdout, "elasticsearch", log.LstdFlags) // 设置日志输出位置
	var err error
	global.ES, err = elastic.NewClient(url, sniff, elastic.SetTraceLog(logger))
	if err != nil {
		zap.S().Panicf("连接es异常: %s", err.Error())
	}

	// 创建mapping
}
