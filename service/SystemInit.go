package service

import (
	"WHisperHArbor-backend/model"
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
)

func ReadConfig() *model.Config {
	co := &model.Config{}
	if configFile, err := ioutil.ReadFile("config.yaml"); err != nil {
		configData, err := yaml.Marshal(co)
		if err != nil {
			panic("failed to marshal default config, err: " + err.Error())
		}

		err = ioutil.WriteFile("config.yaml", configData, 0644)
		if err != nil {
			panic("failed to write default config to file, err: " + err.Error())
		}
		panic("")
	} else {
		if err = yaml.Unmarshal(configFile, co); err != nil {
			panic("failed to unmarshal the config.yaml, err: " + err.Error())
		} else {
			return co
		}
	}
}

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
		err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Favorites{})
		if err != nil {
			panic("failed to migrate the database, err: " + err.Error())
		}
		return db
	}
}
