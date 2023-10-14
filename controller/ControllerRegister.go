package controller

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CheckeAccountExist(account string) (bool, error) {
	var user model.User
	if err := model.DB.Where("account=?", account).First(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	} else {
		return !errors.Is(err, gorm.ErrRecordNotFound), nil
	}

}

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	exist, err := CheckeAccountExist(user.Account)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	if exist {
		c.JSON(http.StatusOK, gin.H{
			"message": "账号已被使用",
			"code":    400,
		})
		return
	}
	if Hashpwd, err := utils.PasswdHash(user.Passwd); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "注册失败",
			"code":    400,
		})
		return
	} else {
		user.Passwd = Hashpwd
	}

	user.AES = utils.RandomString(32)
	if err := model.DB.Create(&user).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"code":    200,
		"user":    user,
	})
}
