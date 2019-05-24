package dao

import (
	"log"
	"qk_video/config"
	"qk_video/lib/dbUtils"
	"qk_video/lib/logger"
	"qk_video/lib/redisUtils"
	"qk_video/module/account/model"
	"qk_video/module/common"
	"time"
)

// AccountDao dao.
type AccountDao struct {
	common.BaseDAo
}

func (d *AccountDao) Init(tableName string) {
	d.Db = dbUtils.CreateConnectionByHost(config.Config.MysqlHost)
	d.TableName = tableName
	d.Redis = redisUtils.InitRedis(config.Config.RedisHost, config.Config.RedisPwd)
}

func (d *AccountDao) AddUser(account *model.Account) error {
	return d.Db.Table(d.TableName).Save(&account).Error
}

func (d *AccountDao) GetCacheCode(key string) string {
	res := d.Redis.Get(key).Val()
	d.Redis.Del(key)
	return res
}

func (d *AccountDao) SetCacheCode(key, val string) {
	if err := d.Redis.Set(key, val, time.Minute*5).Err(); err != nil {
		logger.Err(err)
	}
}

func (d *AccountDao) GetCodeCount(key string) int {
	if d.Redis.Exists(key).Val() == 0 {
		d.Redis.Set(key, 1, time.Hour*24)
		return 1
	}
	res, err := d.Redis.Get(key).Int()
	if err != nil {
		log.Println(err.Error())
		logger.Errorf("GetCodeCount err ", err.Error(), key)
		return 0
	}
	d.Redis.Incr(key)
	return res
}

func (d *AccountDao) ExitsUserByPhone(phone string) bool {
	var count int
	if err := d.Db.Table(d.TableName).Where("phone = ?", phone).Count(&count).Error; err != nil {
		logger.Error(err.Error())
		return false
	}
	return count > 0
}

func (d *AccountDao) GetUserInfoByPwd(phone string, password string) *model.Account {
	var account model.Account
	db := d.Db.Unscoped().
		Table(d.TableName).
		First(&account).
		Where("phone = ? and password = ?", phone, password)

	if db.Error != nil || db.RecordNotFound() {
		logger.Error(db.Error.Error())
		return nil
	}

	return &account
}

func (d *AccountDao) FindUserInfoByID(userID string) *model.Account {
	var account model.Account
	db := d.Db.Unscoped().
		Table(d.TableName).
		First(&account).
		Where("user_id = ? ", userID)

	if db.Error != nil || db.RecordNotFound() {
		logger.Error(db.Error.Error())
		return nil
	}

	return &account
}
