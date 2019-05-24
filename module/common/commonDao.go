package common

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type Dao interface {
	Init(string)
}

type BaseDAo struct {
	Db        *gorm.DB
	Redis     *redis.Client
	TableName string
}
