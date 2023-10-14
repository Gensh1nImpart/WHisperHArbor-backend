package model

import "github.com/golang-jwt/jwt/v4"

type MyClaims struct {
	User LoginUser
	jwt.StandardClaims
}
