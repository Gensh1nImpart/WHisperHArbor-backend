package utils

import (
	"WHisperHArbor-backend/model"
	"github.com/golang-jwt/jwt/v4"
)

func ParseToken(token string) (*model.MyClaims, error) {
	claims := &model.MyClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	return claims, err
}
