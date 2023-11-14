package controller

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/service"
	"WHisperHArbor-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleLogin(userVo model.LoginUser) (bool, string) {
	user := model.User{}
	if err := model.DB.Where("account = ?", userVo.Account).First(&user).Error; err != nil {
		return false, err.Error()
	} else {
		if user.Verify == false {
			go service.LoginWithOutVerify(user.Mail)
			return false, "未验证, 请查看邮箱验证邮件"
		}
		if !utils.PasswdVerify(userVo.Passwd, user.Passwd) {
			return false, "密码错误"
		}
		if token, err := utils.GenerateToken(userVo); err != nil {
			return false, "未知错误"
		} else {
			return true, token
		}
	}
}

func Login(c *gin.Context) {
	var userVo model.LoginUser
	if err := c.ShouldBindJSON(&userVo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "登录失败 " + err.Error(),
		})
		return
	}
	if flag, token := HandleLogin(userVo); flag != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "登录失败 " + token,
			"code":    "400",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "登录成功",
			"token":   token,
		})
	}
}
