package model

import (
	"gorm.io/gorm"
)

var (
	MyConfig = &Config{}
	DB       *gorm.DB
	MyHub    *Hub
)
