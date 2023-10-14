package service

import (
	"WHisperHArbor-backend/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func ReadConfig() *model.Config {
	co := &model.Config{}
	if configFile, err := ioutil.ReadFile("config.yaml"); err != nil {
		panic("failed to read the config.yaml, err: " + err.Error())
	} else {
		if err = yaml.Unmarshal(configFile, co); err != nil {
			panic("failed to unmarshal the config.yaml, err: " + err.Error())
		} else {
			return co
		}
	}
}
