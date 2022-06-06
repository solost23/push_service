package models

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func NewRedisConnect(name string) *redis.Client {
	config := viper.Get("connections.redis." + name).(map[string]interface{})
	if nil == config {
		panic("Redis配置错误")
	}
	addr := fmt.Sprintf("%s:%d", config["host"], config["port"])
	password, ok := config["password"].(string)
	if !ok {
		password = ""
	}
	db := config["db"].(int)
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		Password:   password,
		DB:         db,
		MaxConnAge: 30 * 60 * time.Second,
	})
	return client
}
