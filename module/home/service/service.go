package service

import (
	"qk_video/module/home/dao"
)

var (
	bannerTable = "banner"
)
type HomeService struct {
	dao *dao.HomeDao
}

func NewHomeService() *HomeService {
	d := &dao.HomeDao{}
	d.Init(bannerTable)
	return &HomeService{
		dao: d,
	}
}

func (s *HomeService)GetBannerList() interface{}{
	return s.dao.QueryBannerList()
}

