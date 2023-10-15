package middleware

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func RenweToken(claim *model.MyClaims) (string, error) {
	WithinLimit := func(s, l int64) bool {
		e := time.Now().Unix()
		return e-s < l
	}
	if WithinLimit(claim.ExpiresAt, 600) {
		return utils.GenerateToken(claim.User)
	}
	return "", errors.New("登录已过期")
}

func JWTAuth(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if len(auth) == 0 {
		ctx.Abort()
		ctx.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "无权限",
		})
		return
	}
	if claim, err := utils.ParseToken(auth); err != nil {
		if strings.Contains(err.Error(), "expired") {
			if newToken, _ := RenweToken(claim); newToken != "" {
				ctx.Header("newToken", newToken)
				ctx.Request.Header.Set("Authorization", newToken)
				ctx.Next()
				return
			}
		}
		ctx.Abort()
		ctx.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": err.Error(),
		})
		return
	} else {
		ctx.Next()
	}
}
