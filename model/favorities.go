package model

import "gorm.io/gorm"

type Favorites struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	PostID uint `gorm:"not null"`
}

//func (fa *Favorites) TableName() string {
//	return "favorities"
//}
