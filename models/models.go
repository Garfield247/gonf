package models

import (
	"github.com/Garfield247/gonf.git/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetUp() {
	db, err := gorm.Open(sqlite.Open(config.DataBaseCfg.Path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Services{})
}
