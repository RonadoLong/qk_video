package home

import (
	"github.com/gin-gonic/gin"
	"qk_video/lib/jsonResult"
	"qk_video/module/home/service"
)

var (
	homeService *service.HomeService
)

func InitHomeController()  {
	homeService = service.NewHomeService()
}

func GetBannerList(c *gin.Context) {
	res := homeService.GetBannerList()
	if res == nil {
		jsonResult.CreateNotContent(c)
		return
	}
	jsonResult.CreateSuccess(c, res)
}
