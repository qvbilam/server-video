package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"strconv"
	"video/global"
)

func InitConfig() {
	initViperConfig()
	initEnvConfig()
}

func initEnvConfig() {

	serverPort, _ := strconv.Atoi(os.Getenv("PORT"))
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	esPort, _ := strconv.Atoi(os.Getenv("ES_PORT"))
	redisPort, _ := strconv.Atoi(os.Getenv("REDIS_HOST"))
	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_PASSWORD"))
	userServerPort, _ := strconv.Atoi(os.Getenv("USER-SERVER_PORT"))

	global.ServerConfig.Name = os.Getenv("SERVER_NAME")
	global.ServerConfig.Port = serverPort
	global.ServerConfig.DBConfig.Host = os.Getenv("DB_HOST")

	global.ServerConfig.DBConfig.Port = dbPort
	global.ServerConfig.DBConfig.User = os.Getenv("DB_USER")
	global.ServerConfig.DBConfig.Password = os.Getenv("DB_PASSWORD")
	global.ServerConfig.DBConfig.Database = os.Getenv("DB_DATABASE")

	global.ServerConfig.ESConfig.Host = os.Getenv("ES_HOST")
	global.ServerConfig.ESConfig.Port = esPort

	global.ServerConfig.RedisConfig.Host = os.Getenv("REDIS_HOST")
	global.ServerConfig.RedisConfig.Port = redisPort
	global.ServerConfig.RedisConfig.Password = os.Getenv("REDIS_PASSWORD")
	global.ServerConfig.RedisConfig.Database = redisDb

	global.ServerConfig.UserServerConfig.Name = os.Getenv("USER-SERVER_NAME")
	global.ServerConfig.UserServerConfig.Host = os.Getenv("USER-SERVER_HOST")
	global.ServerConfig.UserServerConfig.Port = int64(userServerPort)
}

func initViperConfig() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("获取配置异常: %s", err)
	}
	// 映射配置文件
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		zap.S().Panicf("加载配置异常: %s", err)
	}
	// 动态监听配置
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&global.ServerConfig)
	})
}
