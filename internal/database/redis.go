package database

import (
	"context"
	"file-system/internal/global"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func init() {
	conf := redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.RedisAddress, global.RedisPort),
		Password: "",
	}
	global.RedisConn = redis.NewClient(&conf)
	_, err := global.RedisConn.Ping(context.TODO()).Result()
	if err != nil {
		fmt.Println("redis ping error")
		return
	}
	fmt.Println("Redis Connection Success ...")
}

func GetRedisConn() *redis.Client {
	return global.RedisConn
}
