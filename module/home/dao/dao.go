package dao

import (
	"qk_video/config"
	"qk_video/lib/dbUtils"
	"qk_video/lib/logger"
	"qk_video/lib/redisUtils"
	"qk_video/module/common"
	"qk_video/module/home/model"
)

type HomeDao struct {
	common.BaseDAo
}

func (d *HomeDao) Init(tableName string) {
	d.Db = dbUtils.CreateConnectionByHost(config.Config.MysqlHost)
	d.TableName = tableName
	d.Redis = redisUtils.InitRedis(config.Config.RedisHost, config.Config.RedisPwd)
}



func (d *HomeDao)QueryBannerList() []model.HomeBanner{
	var bList []model.HomeBanner
	res := d.Db.Table(d.TableName).Where("status = 1").Find(&bList)
	if res.Error != nil || res.RecordNotFound() {
		logger.Errorf("QueryBannerList", res.Error)
		return nil
	}
	return bList
}

