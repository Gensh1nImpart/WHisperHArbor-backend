package service

import (
	"WHisperHArbor-backend/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		model.MyConfig.Mysql.User,
		model.MyConfig.Mysql.Passwd,
		model.MyConfig.Mysql.Host,
		model.MyConfig.Mysql.Port,
		model.MyConfig.Mysql.Database,
		model.MyConfig.Mysql.Charset,
	)

	if db, err := gorm.Open(mysql.Open(args), &gorm.Config{}); err != nil {
		panic("failed to connect the database, err:" + err.Error())
	} else {
		err := db.AutoMigrate(&model.User{})
		if err != nil {
			panic("failed to migrate the database, err: " + err.Error())
		}
		return db
	}
}
