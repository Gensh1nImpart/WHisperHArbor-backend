package utils

import (
	"WHisperHArbor-backend/model"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const TokenExpireDuration = time.Hour * 6

var Secret = []byte("iamshitiloveeatshithhhhasdasdszcarsakjchduiashdi")

func GenerateToken(user model.LoginUser) (string, error) {
	expireDuration := time.Now().Add(TokenExpireDuration)
	claim := &model.MyClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireDuration.Unix(),
			Issuer:    "yrh",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	if tokenString, err := token.SignedString(Secret); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}
