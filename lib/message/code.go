package message

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"wb_stock/common/cache"
	"wb_stock/common/logger"

	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
)

const (
	appkey          = "8af596c745e7ce9eedd7bcea54fb114c"
	content         = "【TOH】您的验证码是"
	contentRegister = "【TOH僑网】您注册的用户名是%s,密码是%s,祝您投资愉快"
)

func SendCode(phone string) string {
	rand.Seed(time.Now().Unix())
	code := strconv.Itoa(rand.Intn(1000000))
	cache.SetCodeToCache(phone, code)
	time.Now()
	clnt := ypclnt.New(appkey)
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = phone
	param[ypclnt.TEXT] = content + code
	r := clnt.Sms().SingleSend(param)
	logger.Info(r)
	if r.Code == 0 {
		return code
	}
	logger.Info("get code error")
	return ""
}

func SendRegisterMsg(phone, password, username string) error {
	clnt := ypclnt.New(appkey)
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = phone
	con := fmt.Sprintf(contentRegister, username, password)
	logger.Info(con)
	param[ypclnt.TEXT] = con
	r := clnt.Sms().BatchSend(param)
	logger.Info(r)
	if r.Code == 0 {
		return nil
	}
	logger.Info("get code error")
	return errors.New(r.Msg)
	//账户:clnt.User() 签名:clnt.Sign() 模版:clnt.Tpl() 短信:clnt.Sms() 语音:clnt.Voice() 流量:clnt.Flow()
}
