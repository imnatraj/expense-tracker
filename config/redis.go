package config

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisDBType int

const (
	RedisDBTypeOTP     RedisDBType = 0
	RedisDBTypeSession RedisDBType = 1
)

func GetRedisDbClient(host, password string, variant RedisDBType) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", host),
		Password: password,     // no password set
		DB:       int(variant), // use default DB
	})
	return rdb
}
