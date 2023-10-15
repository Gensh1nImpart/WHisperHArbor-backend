package model

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Content   string    `gorm:"type:text;not null" json:"content"`
	Time      time.Time `gorm:"type:datetime;not null"`
	Encrypted bool      `gorm:"type:boolean;not null;default: false" json:"encrypted"`
	UserId    uint      `json:"userId"`
	Likes     int64     `gorm:"type:int;default:0"`
	User      User      `gorm:"foreignKey: UserId"`
}

//func (post *Post) TableName() string {
//	return "Post"
//}

type AddLikes struct {
	PostId uint `json:"post_id"`
}
