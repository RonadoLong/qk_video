package middleware

import (
	"github.com/gin-gonic/gin"
	"qk_video/lib/jsonResult"
	"qk_video/lib/jwt"
)

const Authorization = "Authorization"

// AuthMiddleware token middleware
func AuthMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			jsonResult.CreateErrorWithMsg(c, "请求未携带token，无权限访问")
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				jsonResult.CreateErrorWithMsg(c, "授权已过期")
				return
			}
			jsonResult.CreateErrorWithMsg(c, err.Error())
			return
		}
		c.Set("role", claims.VipLevel)
		c.Set("id", claims.UserID)
	}
}
