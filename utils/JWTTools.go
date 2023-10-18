package utils

import (
	"WHisperHArbor-backend/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func ParseToken(token string) (*model.MyClaims, error) {
	claims := &model.MyClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	return claims, err
}

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
