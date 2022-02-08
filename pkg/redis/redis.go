package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var RedisClient *redis.Client

func SetUp() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "114.117.161.250:9380",
		Password: "blackb", // no password set
		DB:       0,        // use default DB
	})
	err := RedisClient.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Panicln(err)
	}
}

func Set(key string, data interface{}, time int) error {
	err := RedisClient.Set(ctx, "key", "value", 0).Err()
	return err
}

func Get(key string) (string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val, err
}
