package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Account  string `gorm:"type:varchar(20);not null" json:"account"`
	Passwd   string `gorm:"type:varchar(125);not null" json:"passwd"`
	Mail     string `gorm:"type:varchar(125);not null" json:"mail"`
	Verify   bool   `json:"verify"`
	Nickname string `gorm:"type:varchar(20);not null" json:"nickname"`
	AES      string `gorm:"type:varchar(125);not null" `
}

//func (user *User) TableName() string {
//	return "User"
//}

type LoginUser struct {
	Account string `json:"account"`
	Passwd  string `json:"passwd"`
}
