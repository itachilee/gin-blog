package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var RedisClient *redis.Client

func SetUp() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "114.117.161.250:9280",
		Password: "blackb", // no password set
		DB:       0,        // use default DB
	})

	err := Set("key", "value", 0)
	if err != nil {
		panic(err)
	}

	val, err := Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	log.Println("redis test ok")
}

func Set(key string, data interface{}, time time.Duration) error {
	err := RedisClient.Set(ctx, "key", "value", time).Err()
	return err
}

func Get(key string) (string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val, err
}
