package env

import (
	"qk_video/module/account"
	"qk_video/module/home"
)

func InitAllEnv()  {
	account.InitUserEnv()
	home.InitHomeController()
}


