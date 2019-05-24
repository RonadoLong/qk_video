package common

import (
	"github.com/gin-gonic/gin"
	"qk_video/lib/codeCaptchaUtils"
	"qk_video/lib/jsonResult"
)

func GetImageCode(c *gin.Context) {
	codeID, imageStr := codeCaptchaUtils.CodeCaptchaCreate()
	jsonResult.CreateSuccess(c, gin.H{
		"codeID": codeID,
		"imageStr": imageStr,
	})
}
