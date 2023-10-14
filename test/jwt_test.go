package test

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/utils"
	"log"
	"testing"
)

func TestJWT(t *testing.T) {
	log.Println(utils.GenerateToken(model.LoginUser{
		Account: "test",
		Passwd:  "test",
	}))
}
