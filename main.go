package main

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/service"
	"github.com/gin-gonic/gin"
)

func init() {
	model.MyConfig = service.ReadConfig()
	model.DB = service.InitDB()
}

func main() {
	r := gin.Default()
	MyRouter(r)
	r.Run(":" + model.MyConfig.Base.Port)
}
