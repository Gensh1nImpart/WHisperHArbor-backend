package controller

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/service"
	"WHisperHArbor-backend/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CheckeAccountExist(account string, mail string) (bool, error) {
	var user model.User
	if err := model.DB.Where("account=?", account).First(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		if err = model.DB.Where("mail=?", mail).First(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		} else {
			return !errors.Is(err, gorm.ErrRecordNotFound), nil
		}
	} else {
		return !errors.Is(err, gorm.ErrRecordNotFound), nil
	}

}

func Register(c *gin.Context) {
	var temp struct {
		Account  string `json:"account"`
		Passwd   string `json:"passwd"`
		Mail     string `json:"mail"`
		Nickname string `json:"nickname"`
	}
	if err := c.ShouldBindJSON(&temp); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "格式有误!",
		})
		return
	}
	user := &model.User{
		Account:  temp.Account,
		Passwd:   temp.Passwd,
		Mail:     temp.Mail,
		Verify:   false,
		Nickname: temp.Nickname,
	}
	exist, err := CheckeAccountExist(user.Account, user.Mail)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "注册失败!",
		})
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
	go service.SendEmailForVerify(user.Mail)
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功, 请前去验证",
		"code":    200,
		"user":    user,
	})
}
