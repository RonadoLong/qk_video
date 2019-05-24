package router

import (
	"log"
	"net/http"
	"qk_video/config"
	"qk_video/lib/middleware"
	"qk_video/module/account"
	"qk_video/module/common"
	"qk_video/module/home"
	"time"

	"github.com/gin-gonic/gin"
)

// StartWebServer 开启web服务
func StartWebServer() {
	router := gin.Default()

	api := router.Group("api/client")
	{
		//api.Use(middleware.Cors())
		api.Use(middleware.LimiterMiddleware())
		api.Use(middleware.LogoInterceptor())
		api.GET("common/get_image_code", common.GetImageCode)
	}

	accountApi := api.Group("/account")
	{
		accountApi.POST("login",  account.Login)
		accountApi.POST("register", account.Register)
		accountApi.GET("get_code/:phone", account.GetCode)
		accountApi.GET("info",middleware.AuthMiddleware(), account.FindUserInfo)
	}

	homeApi := api.Group("/home")
	{
		homeApi.GET("/banners", home.GetBannerList)
	}

	port := config.Config.ServerHost
	server := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening server at ", port)
	_ = server.ListenAndServe()
}
