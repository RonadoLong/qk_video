package config

import (
	"fmt"
	"log"
	"qk_video/lib/logger"

	"github.com/spf13/viper"
)

const (
	tokenKeyName      = "token.%s.key"
	serverHostName    = "server.%s.host"
	mysqlHostName     = "mysql.%s.host"
	mysqlLogLevelName = "mysql.%s.logLevel"
	redisHostName     = "redis.%s.host"
	redisPwdName      = "redis.%s.password"
)

var Config *conf

type conf struct {
	TokenKey    string
	ServerHost  string
	HystrixPort string

	MysqlHost     string
	MysqlLogLevel string

	RedisHost string
	RedisPwd  string
}

const (
	path = "config"
)

// Init Init env setting
func Init(env string) {
	var err error
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath(path)

	err = v.ReadInConfig()
	if err != nil {
		log.Println("error on parsing message", err.Error())
	}
	logger.Infof("当前配置是 ======= ", env)
	Config = &conf{}

	Config.MysqlHost = v.GetString(fmt.Sprintf(mysqlHostName, env))
	Config.MysqlLogLevel = v.GetString(fmt.Sprintf(mysqlLogLevelName, env))

	Config.TokenKey = v.GetString(fmt.Sprintf(tokenKeyName, env))
	Config.ServerHost = v.GetString(fmt.Sprintf(serverHostName, env))

	Config.RedisHost = v.GetString(fmt.Sprintf(redisHostName, env))
	Config.RedisPwd = v.GetString(fmt.Sprintf(redisPwdName, env))
}
