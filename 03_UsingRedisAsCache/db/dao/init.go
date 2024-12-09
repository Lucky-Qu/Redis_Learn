package dao

import (
	"Redis_Learn/03_UsingRedisAsCache/config"
	"Redis_Learn/03_UsingRedisAsCache/db/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: config.DSN,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			TablePrefix:   "UsingRedisAsCache_",
		},
	})
	if err != nil {
		panic(err)
	}
	if !db.Migrator().HasTable(&model.User{}) {
		err = db.Migrator().CreateTable(&model.User{})
		if err != nil {
			panic(err)
		}
	}
}
