package main

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func init() {
	model.MyConfig = service.ReadConfig()
	model.DB = service.InitDB()
	model.MyHub = model.NewHub()
	go model.MyHub.Run()
}

func main() {
	r := gin.Default()
	now := time.Now()
	dateString := now.Format("2006-01-02")
	logFileName := fmt.Sprintf("%s/%v-log.log", "logs", dateString)
	f, err := os.Create(logFileName)
	if err != nil {
		panic("can create log file")
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	MyRouter(r)
	r.Run(":" + model.MyConfig.Base.Port)
}
