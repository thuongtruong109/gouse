package connection

import "github.com/go-redis/redis/v8"

func Redis(addr, pass string, dbNo ...int) *redis.Client {
	nums := 0

	if len(dbNo) > 0 {
		nums = dbNo[0]
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       nums,
	})
	return rdb
}

// import (
// 	"github.com/go-redis/redis"
// 	"github.com/spf13/viper"
// )

// var RedisClient *redis.Client

// func InitRedisClient() {
// 	RedisClient = redis.NewClient(&redis.Options{
// 		Addr:     viper.GetString("redis.host"),
// 		Password: viper.GetString("redis.password"),
// 		DB:       viper.GetInt("redis.db"),
// 	})
// }

// func CloseRedisClient() {
// 	_ = RedisClient.Close()
// }
