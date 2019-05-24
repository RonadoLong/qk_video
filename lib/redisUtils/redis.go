package redisUtils

import (
	"github.com/go-redis/redis"
	"os"
	"qk_video/lib/logger"
	"time"
)

//声明一些全局变量
func InitRedis(host string, password string) *redis.Client{
	redisClient := redis.NewClient(&redis.Options{
		Addr:        host,
		Password:    password,
		DB:          0, // use default DB
		IdleTimeout: 240 * time.Second,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		logger.Errorf("connec redis err " + err.Error())
		os.Exit(1)
		return nil
	}
	//go func() {
	//	sub := redisClient.PSubscribe("__keyevent@0__:expired")
	//	_, _ = sub.ReceiveTimeout(time.Second * 110)
	//	for message := range sub.Channel() {
	//		logger.Infof("__keyevent === expired , %s", message.Payload)
	//	}
	//}()
	return redisClient
}
