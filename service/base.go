package service

import (
	"github.com/AnnatarHe-Athena/hookman/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	d, err := gorm.Open("postgres", config.DB_DSN)
	if err != nil {
		panic(err)
	}

	d.LogMode(true)

	db = d
}
