package main

import (
	"WHisperHArbor-backend/controller"
	"WHisperHArbor-backend/middleware"
	"WHisperHArbor-backend/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func MyRouter(r *gin.Engine) {
	authapi := r.Group("/auth")
	{
		authapi.POST("/login", controller.Login)
		authapi.POST("/register", controller.Register)
	}
	authorized := r.Group("/v1/api", middleware.JWTAuth)
	{
		authorized.GET("/sayhello", func(c *gin.Context) {
			auth := c.Request.Header.Get("Authorization")
			claims, _ := utils.ParseToken(auth)
			log.Println(claims)
			c.String(http.StatusOK, "hello"+claims.User.Account)
		})
	}
}
