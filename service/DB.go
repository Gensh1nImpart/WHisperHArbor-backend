package service

import (
	"WHisperHArbor-backend/model"
)

//func checkLogin(user model.LoginUser) (bool, error) {
//
//}

func GetUser(claim model.MyClaims) (model.User, error) {
	user := model.User{}
	if err := model.DB.Where("account = ?", claim.User.Account).First(&user).Error; err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func UserPost(claim model.User) ([]model.Post, error) {
	var posts []model.Post
	if err := model.DB.Where("user_id = ?", claim.ID).Find(&posts).Error; err != nil {
		return posts, err
	} else {
		return posts, nil
	}
}

func PublicPost() ([]model.Post, error) {
	var posts []model.Post
	if err := model.DB.Preload("User").Where("encrypted = ?", false).Find(&posts).Error; err != nil {
		return posts, nil
	} else {
		return posts, nil
	}
}
