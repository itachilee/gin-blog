package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	setting "github.com/itachilee/ginblog/internal/setting"
)

var ctx = context.Background()

var Client *redis.Client

func InitRedis() {
	fmt.Println(setting.RedisSetting.Password)
	Client = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Addr,
		Username: setting.RedisSetting.Username,
		Password: setting.RedisSetting.Password, // no password set
		DB:       0,                             // use default DB
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
