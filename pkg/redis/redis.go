package redis

import (
	"context"
	"fmt"
	"github.com/itachilee/ginblog/global"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var Client *redis.Client

func InitRedis() {
	fmt.Println(global.RedisSetting.Password)
	Client = redis.NewClient(&redis.Options{
		Addr:     global.RedisSetting.Addr,
		Username: global.RedisSetting.Username,
		Password: global.RedisSetting.Password, // no password set
		DB:       0,                            // use default DB
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
	err := Client.Set(ctx, "key", "value", time).Err()
	return err
}

func Get(key string) (string, error) {
	val, err := Client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val, err
}
