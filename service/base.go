package service

import (
	"github.com/AnnatarHe-Athena/hookman/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	d, err := gorm.Open(postgres.Open(config.DB_DSN), &gorm.Config{})
	if err != nil {
		logrus.Panicln(err)
	}

	db = d
}
