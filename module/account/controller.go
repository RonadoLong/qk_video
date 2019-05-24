package account

import (
	"github.com/gin-gonic/gin"
	"log"
	"qk_video/lib"
	"qk_video/lib/codeCaptchaUtils"
	"qk_video/lib/jsonResult"
	"qk_video/lib/jwt"
	"qk_video/lib/logger"
	"qk_video/module/account/model"
	"qk_video/module/account/service"
	"strings"
)

var (
	srv *service.UserService
)

func InitUserEnv() {
	srv = service.NewUserService()
}

func Register(ctx *gin.Context) {
	var req model.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Errorf("bind data err ", err.Error())
		jsonResult.CreateError(ctx)
		return
	}

	var phone = req.Phone
	var userID = lib.GetUUID()
	var username = lib.StringJoinString("Nice", phone)

	var code = srv.GetCode(phone)
	if code == "" {
		jsonResult.CreateErrorWithMsg(ctx, lib.StringJoinString(phone, " 验证码不存在 "))
		return
	}

	if !strings.EqualFold(strings.TrimSpace(code), strings.TrimSpace(req.Code)) {
		jsonResult.CreateErrorWithMsg(ctx, lib.StringJoinString(phone, " 验证码错误 "))
		return
	}

	if exits := srv.ExitsUserByPhone(phone); exits {
		jsonResult.CreateErrorWithMsg(ctx, lib.StringJoinString(phone, " 已注册，请登录"))
		return
	}

	account := &model.Account{
		UserID:    userID,
		Username:  username,
		Phone:     phone,
		Password:  req.Password,
		Recommend: req.Recommend,
		VipLevel:  0,
	}

	if err := srv.Register(account); err != nil {
		logger.Errorf("Register err ", err)
		jsonResult.CreateError(ctx)
		return
	}

	option := &jwt.Option{
		UserID:   userID,
		Name:     username,
		VipLevel: 0,
		Phone:    phone,
	}
	token, err := jwt.CreateToken(jwt.NewCustomClaims(option))
	if err != nil {
		log.Println(err)
		jsonResult.CreateError(ctx)
		return
	}

	accountResp := model.AccountResp{
		UserID: userID,
		Token: token,
		Level: 0,
		Username: username,
	}
	jsonResult.CreateSuccess(ctx, accountResp)
}

func FindUserInfo(c *gin.Context) {
	userID := c.GetString("id")
	if userID == "" {
		return
	}

	account := srv.FindUserInfoByID(userID)
	if account == nil {
		jsonResult.CreateErrorWithMsg(c, "user not exits")
		return
	}

	jsonResult.CreateSuccess(c, account)
}

func Login(c *gin.Context) {
	var req model.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Infof("bind data err ", err.Error())
		jsonResult.CreateErrorParams(c)
		return
	}

	res := codeCaptchaUtils.VerifyCaptcha(req.CodeID, req.VerifyVal)
	if !res {
		jsonResult.CreateErrorWithMsg(c, "验证码有误")
		return
	}

	account := srv.GetUserInfo(req.Password, req.Password)
	if account == nil {
		jsonResult.CreateErrorWithMsg(c, "phone or password err")
		return
	}
	option := &jwt.Option{
		UserID:   account.UserID,
		Name:     account.Username,
		VipLevel: account.VipLevel,
		Phone:    account.Phone,
	}
	token, err := jwt.CreateToken(jwt.NewCustomClaims(option))
	if err != nil {
		log.Println(err)
		jsonResult.CreateError(c)
		return
	}

	accountResp := model.AccountResp{
		UserID: account.UserID,
		Token: token,
		Level: account.VipLevel,
		Username: account.Username,
	}
	jsonResult.CreateSuccess(c, accountResp)
}

func GetCode(ctx *gin.Context) {
	phone := ctx.Param("phone")
	if phone == "" || !lib.CheckMobileNum(phone) {
		jsonResult.CreateErrorParams(ctx)
		return
	}
	isMaxGetCode := srv.IsMaxGetCode(phone)
	if isMaxGetCode {
		jsonResult.CreateErrorWithMsg(ctx, "获取code已达最大次数，请明天再试")
		return
	}
	code := lib.GetRandCode()
	srv.SetCode(phone, code)
	jsonResult.CreateSuccess(ctx, code)
}
