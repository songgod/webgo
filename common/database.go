package common

import (
	"back/model"

	_ "github.com/glebarez/sqlite"
	"github.com/jinzhu/gorm"
)

var singledb *gorm.DB

func InitDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite", "user.db")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})

	singledb = db

	return
}

func GetDB() *gorm.DB {
	return singledb
}
