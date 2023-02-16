package common

import (
	"back/model"

	"github.com/jinzhu/gorm"
)

var singledb *gorm.DB

func InitDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", "user.db")
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
