package controller

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

/*
@ 获取Websocket，http升级为长轮寻
@ method: GET
@ Headers: Authorization
*/
func ServeWS(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	claim, _ := utils.ParseToken(auth)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "连接失败!" + err.Error(),
		})
		return
	}
	Client := &model.Client{Hub: model.MyHub, Conn: conn, Send: make(chan model.Message), User: []byte(claim.User.Account)}
	Client.Hub.Register <- Client
	go Client.WritePump()
	go Client.ReadPump()
}
