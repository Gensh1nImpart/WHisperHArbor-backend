package test

import (
	"WHisperHArbor-backend/utils"
	"log"
	"testing"
)

func TestPass(t *testing.T) {
	bytes, _ := utils.PasswdHash("123123")
	log.Println(utils.PasswdVerify("123123", bytes))
}
