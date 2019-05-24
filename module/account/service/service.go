package service

import (
	"github.com/pkg/errors"
	"qk_video/lib"
	"qk_video/module/account/dao"
	"qk_video/module/account/model"
)

const (
	phoneLimiter = ":limiter"
	phoneCode = ":code"
)

var (
	commonErr =  errors.New("服务器保存失败")
	accountTable = "account"
)

type UserService struct {
	dao *dao.AccountDao
}

func NewUserService() *UserService{
	d := new(dao.AccountDao)
	d.Init(accountTable)
	return &UserService{
		dao: d,
	}
}

func (s *UserService)GetUserInfo(phone, password string) *model.Account{
	return s.dao.GetUserInfoByPwd(phone, password)
}

func (s *UserService) FindUserInfoByID(userID string) *model.Account{
	return s.dao.FindUserInfoByID(userID)
}

func (s *UserService)ExitsUserByPhone(phone string) bool {
	return s.dao.ExitsUserByPhone(phone)
}

// redis
func (s *UserService)GetCode(key string) string {
	key = lib.StringJoinString(key, phoneCode)
	return s.dao.GetCacheCode(key)
}

func (s *UserService)SetCode(key, val string) {
	key = lib.StringJoinString(key, phoneCode)
	s.dao.SetCacheCode(key, val)
}

func (s *UserService)IsMaxGetCode(key string) bool {
	key = lib.StringJoinString(key, phoneLimiter)
	codeCount := s.dao.GetCodeCount(key)
	if codeCount == 0|| codeCount >= 5 {
		return true
	}
	return false
}

func (s *UserService)Register(account *model.Account) interface{} {
	if err := s.dao.AddUser(account); err != nil {
		return commonErr
	}
	return nil
}


