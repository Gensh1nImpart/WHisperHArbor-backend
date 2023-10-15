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
	User      User      `gorm:"foreignKey: UserId"`
}
