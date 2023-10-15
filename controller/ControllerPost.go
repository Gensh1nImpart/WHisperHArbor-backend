package controller

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/service"
	"WHisperHArbor-backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func UserPost(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	claim, _ := utils.ParseToken(auth)
	PostContext := model.Post{}
	if err := c.ShouldBindJSON(&PostContext); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "发布失败!" + err.Error(),
		})
		return
	} else {
		if user, err := service.GetUser(*claim); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "发布失败!" + err.Error(),
			})
			return
		} else {
			if PostContext.Encrypted {
				content, err := utils.EncryptPost(PostContext.Content, user.AES)
				if err != nil {
					c.JSON(http.StatusOK, gin.H{
						"code":    400,
						"message": "发布失败!" + err.Error(),
					})
					return
				}
				PostContext.Content = content
			}
			PostContext.UserId = user.ID
			PostContext.Time = time.Unix(time.Now().Unix(), 0)
			if err := HandlePost(&PostContext); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    400,
					"message": "发布失败!" + err.Error(),
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": "发布成功!",
				})
			}
		}
	}
}

func HandlePost(post *model.Post) error {
	if err := model.DB.Create(post).Error; err != nil {
		return err
	} else {
		return nil
	}
}
