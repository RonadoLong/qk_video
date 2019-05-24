package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	VipLevel int8   `json:"vipLevel"`
	jwt.StandardClaims
}

type Option struct {
	UserID   string
	Name     string
	Phone    string
	VipLevel int8
}

func NewCustomClaims(option *Option) *CustomClaims {

	c := &CustomClaims{
		UserID:   option.UserID,
		Name:     option.Name,
		Phone:    option.Phone,
		VipLevel: option.VipLevel,
	}
	c.ExpiresAt = time.Now().Add(expireTime).Unix()
	c.Issuer = string(SignKey)
	return c
}
