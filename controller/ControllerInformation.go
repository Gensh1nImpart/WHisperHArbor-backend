package controller

import (
	"WHisperHArbor-backend/service"
	"WHisperHArbor-backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OldNew struct {
	Old string `json:"old"`
	New string `json:"New"`
}

func ModifyPasswd(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	claim, _ := utils.ParseToken(auth)
	temp := &OldNew{}
	if err := c.ShouldBindJSON(&temp); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "修改失败" + err.Error(),
		})
		return
	} else {
		if claim.User.Passwd != temp.Old {
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "修改失败" + err.Error(),
			})
			return
		}
		if hashwd, err := utils.PasswdHash(temp.New); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "修改失败" + err.Error(),
			})
			return
		} else {
			flag, err := service.ModifyPasswd(*claim, hashwd)
			if flag {
				c.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": "修改成功",
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    400,
					"message": "修改失败" + err.Error(),
				})
				return
			}
		}
	}
}
