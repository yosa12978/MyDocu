package db

import (
	"fmt"
	"sync"

	"github.com/yosa12978/MyDocu/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func GetDB() *gorm.DB {
	once.Do(func() {
		var err error
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			config.Conf.Db.DbHost,
			config.Conf.Db.Username,
			config.Conf.Db.Password,
			config.Conf.Db.DbName,
			config.Conf.Db.DbPort,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
	return db
}
