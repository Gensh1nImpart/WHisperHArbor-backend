package main

import (
	"WHisperHArbor-backend/controller"
	"WHisperHArbor-backend/middleware"
	"WHisperHArbor-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
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
			if claims, err := utils.ParseToken(auth); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 400,
					"msg":  "token error",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "hello" + claims.User.Account,
				})
			}
		})
		authorized.GET("/personalPost", controller.UserGetPost)
		authorized.GET("/publicPost", controller.PublicGetPost)
		authorized.GET("/personalFavorites", controller.GetUserFavorites)
		authorized.POST("/PostLikes", controller.UserLikePost)
		authorized.POST("/FavoritePost", controller.UserFavoritePost)
		authorized.POST("/post", controller.UserPost)
		authorized.GET("/ws", controller.ServeWS)
	}
}
