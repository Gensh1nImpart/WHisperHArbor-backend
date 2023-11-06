package controller

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServeWS(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	claim, _ := utils.ParseToken(auth)
	conn, err := model.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "连接失败!" + err.Error(),
		})
		return
	}
	Client := &model.Client{Hub: model.MyHub, Conn: conn, Send: make(chan model.Message), User: &claim.User}
	Client.Hub.Register <- Client
	go Client.WritePump()
	go Client.ReadPump()
}
