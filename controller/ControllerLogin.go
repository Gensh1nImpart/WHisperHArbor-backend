package controller

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleLogin(userVo model.LoginUser) (bool, string) {
	user := model.User{}
	if err := model.DB.Where("account = ?", userVo.Account).First(&user).Error; err != nil {
		return false, ""
	} else {
		if !utils.PasswdVerify(userVo.Passwd, user.Passwd) {
			return false, ""
		}
		if token, err := utils.GenerateToken(userVo); err != nil {
			return false, ""
		} else {
			return true, token
		}
	}
}

func Login(c *gin.Context) {
	var userVo model.LoginUser
	if err := c.ShouldBindJSON(&userVo); err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	if flag, token := HandleLogin(userVo); flag != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "登录失败",
			"code":    "400",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "登录成功",
			"token":   token,
		})
	}
	/*
		Just For Test
		var db = &model.LoginUser{Account: "test", Passwd: "123321"}
		if userVo.Account == db.Account && userVo.Passwd == db.Passwd {
			token, _ := utils.GenerateToken(userVo)
			c.JSON(http.StatusOK, gin.H{
				"code": 201,
				"jwt":  token,
				"msg":  "登录成功",
			})
			return
		} else {
			c.String(http.StatusOK, "error")
		}
		return
	*/
}
