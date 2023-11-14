package service

import (
	"WHisperHArbor-backend/model"
	"WHisperHArbor-backend/utils"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

var ctx = context.Background()

func InitializeStore() *model.StorageService {
	myDB, err := strconv.ParseInt(model.MyConfig.Redis.DB, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     model.MyConfig.Redis.Addr,
		Password: model.MyConfig.Redis.Passwd,
		DB:       int(myDB),
	})
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	return &model.StorageService{
		RedisClient: redisClient,
	}
}

func LoginWithOutVerify(mail string) {
	_, err := model.MyRedis.RedisClient.Get(ctx, mail).Result()
	if err != nil {
		SendEmailForVerify(mail)
	}
}

func SendEmailForVerify(mail string) {
	timeStamp := time.Now().Unix()
	data := []byte(fmt.Sprintf("%s%d", mail, timeStamp))
	hash := sha256.Sum256(data)
	key := hex.EncodeToString(hash[:32])
	err := SetKV(key, mail)
	if err == nil {
		_ = utils.SendMail([]string{mail}, "http://localhost:8000/auth/verify/"+key)
	}
}

func SetKV(key string, value string) error {
	err := model.MyRedis.RedisClient.Set(ctx, key, value, time.Hour*24).Err()
	if err != nil {
		return err
	}
	err = model.MyRedis.RedisClient.Set(ctx, value, 1, time.Hour*24).Err()
	return err
}

func GetKV(key string) (string, error) {
	value, err := model.MyRedis.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	err = model.MyRedis.RedisClient.Del(ctx, key).Err()
	if err != nil {
		return value, err
	}
	_ = model.MyRedis.RedisClient.Del(ctx, value).Err()
	return value, err
}
