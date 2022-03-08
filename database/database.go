package database

import (
	"github.com/d0kur0/cui-server/graph/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (err error) {
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.UserToken{},
		&model.Service{},
	)

	if err != nil {
		return
	}

	return
}

func GetDB() *gorm.DB {
	return db
}
