package main

import (
	"WHisperHArbor-backend/controller"
	"github.com/gin-gonic/gin"
)

func MyRouter(r *gin.Engine) {
	authapi := r.Group("/auth")
	{
		authapi.POST("/login", controller.Login)
		authapi.POST("/register", controller.Register)
	}
}
