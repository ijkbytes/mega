package service

import (
	"fmt"
	"time"

	"github.com/ijkbytes/mega/base/config"
	"github.com/ijkbytes/mega/base/log"
	"github.com/ijkbytes/mega/model"
	"github.com/jinzhu/gorm"
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

	db.LogMode(config.Mega.Db.Debug)

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(5 * time.Minute)

	if config.Mega.Db.Migrate {
		db := db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")
		if err := db.AutoMigrate(model.AllModels...).Error; err != nil {
			logger.Fatal("auto migrate tables failed: " + err.Error())
		}
	}
}

func DisconnectDB() {
	if db == nil {
		return
	}
	if err := db.Close(); err != nil {
		logger.Sugar().Errorf("Disconnect from database failed: " + err.Error())
	}
}
