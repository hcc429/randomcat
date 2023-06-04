package db

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
	redisAddr  = os.Getenv("REDIS_ADDR")
	useCache = os.Getenv("USE_CACHE") == "TRUE"
)

func init() {
	if !useCache{
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr, // Update with the Redis server address
		Password: "",         // No password set
		DB:       0,          // Use the default DB
	})
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected to Redis!")
	}
	
}

func GetValue(key string) (string, error) {
	if !useCache{
		return "", errors.New("Cache is disable")
	}
	ctx := context.Background()
	val, err := RedisClient.Get(ctx, key).Result()
	return val, err
}

func AddKeyValuePair(key string, val interface{}, expireTime int) {
	if !useCache{
		return
	}
	ctx := context.Background()
	err := RedisClient.Set(ctx, key, val, time.Duration(expireTime)*time.Second).Err()
	if err != nil {
		panic(err)
	}
}
