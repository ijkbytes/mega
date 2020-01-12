package service

import (
	"fmt"
	"github.com/ijkbytes/mega/base/config"
	"github.com/ijkbytes/mega/base/log"
	"github.com/jinzhu/gorm"
	"time"
)

var logger = log.Get("service")

var db *gorm.DB

func ConnectDB() {
	var err error
	db, err = gorm.Open(
		config.Mega.Db.Type,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Mega.Db.User,
			config.Mega.Db.Password,
			config.Mega.Db.Host,
			config.Mega.Db.Name,
		),
	)
	if err != nil {
		logger.Sugar().Panic("init db err: ", err)
		return
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.Mega.Db.TablePrefix + defaultTableName
	}

	//db.LogMode(true)

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(5 * time.Minute)

	//db.CreateTable(&model.Tag{})
}

func DisconnectDB() {
	if db == nil {
		return
	}
	if err := db.Close(); err != nil {
		logger.Sugar().Errorf("Disconnect from database failed: " + err.Error())
	}
}
