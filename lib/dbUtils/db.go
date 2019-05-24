package dbUtils

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// CreateConnectionByHost 建立连接
func CreateConnectionByHost(host string) *gorm.DB {
	if host == "" {
		log.Panic("host not allow null")
	}

	var err error
	var db *gorm.DB

	db, err = gorm.Open("mysql", host)
	if nil != err {
		log.Panic("opens database failed: " + err.Error())
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	go func() {
		for {
			<-time.After(time.Minute * 2)
			_ = db.DB().Ping()
		}
	}()
	return db
}
