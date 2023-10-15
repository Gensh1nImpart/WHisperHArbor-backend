package service

import "WHisperHArbor-backend/model"

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
