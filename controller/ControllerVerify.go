package controller

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleVerify(c *gin.Context) {
	key := c.Param("key")
	if value, err := service.GetKV(key); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	} else {
		if err = model.DB.Where("mail=?", value).First(&model.User{}).Update("verify", true).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "successful verify",
			})
		}
	}
}
