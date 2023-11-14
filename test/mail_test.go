package test

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/service"
	"WHisperHArbor-backend/utils"
	"log"
	"sync"
	"testing"
)

func TestMail(t *testing.T) {
	model.MyMail = service.ReadEmail()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		err := utils.SendMail([]string{"1223605525@qq.com"}, "http://2Flocalhost%3A8000%2Fverify%2Fqweoijsada&subtemplate=gray&evil=0localhost:8000/verify/qweoijsada")
		log.Println(err)
		wg.Done()
	}()
	wg.Wait()
}

func TestHash(t *testing.T) {
	service.SendEmailForVerify("1@myyrh.com")
}
