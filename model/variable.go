package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type StorageService struct {
	RedisClient *redis.Client
}

var (
	MyConfig = &Config{}
	DB       *gorm.DB
	MyHub    *Hub
	MyRedis  *StorageService
	MyMail   *Email
)
