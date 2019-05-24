package main

import (
	"flag"
	"qk_video/config"
	"qk_video/module/env"
	"qk_video/module/router"

	"github.com/gin-gonic/gin"
)

func init() {
	envStr := flag.String("e", "dev", "环境变量")
	flag.Parse()
	if *envStr == "" || envStr == nil {
		flag.Usage()
	}
	config.Init(*envStr)

	if *envStr == "pro" {
		gin.SetMode(gin.ReleaseMode)
	}
	env.InitAllEnv()
}

func main() {
	router.StartWebServer()
}
