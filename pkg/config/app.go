package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "go_test:password@tcp(127.0.0.1:3306)/go_database?charset=utf8&parseTime=True")

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDb() *gorm.DB {
	return db
}
