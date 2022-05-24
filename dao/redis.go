package dao

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var (
	rdb *redis.Client
)

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "bq134721", // no password set
		DB:       0,          // use default DB
	})
}

// redis设置值
func RedisSet(key string, value string) bool {
	err := rdb.Set(ctx, key, value, 0).Err()

	if err != nil {
		return false
	} else {
		return true
	}
}

func RedisGet(key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		return ""
	} else {
		return val
	}
}
