package service

import (
	"github.com/AnnatarHe-Athena/hookman/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	gormlog "github.com/onrik/logrus/gorm"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func init() {
	d, err := gorm.Open("postgres", config.DB_DSN)
	if err != nil {
		logrus.Panicln(err)
	}

	d.SetLogger(gormlog.New(logrus.StandardLogger()))
	d.LogMode(true)
	d.Set("gorm:table_options", "charset=utf8")

	db = d
}
