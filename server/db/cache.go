package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
	redis_Addr  = os.Getenv("REDIS_ADDR")
	useCache = os.Getenv("USE_REDIS")
)

func init() {
	//Redis
	fmt.Println("shit")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redis_Addr, // Update with the Redis server address
		Password: "",         // No password set
		DB:       0,          // Use the default DB
	})
	pong, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Println(pong, err)
		panic("sad")
	}
}

func GetValue(s string) (string, error) {
	if(useCache!="TRUE"){
		return "No Cache", nil
	}
	ctx := context.Background()
	val, err := RedisClient.Get(ctx, s).Result() // => GET key2
	if err == redis.Nil {
		//log.Println("key2 does not exist")
		return "", errors.New("Redis no data")
	} else if err != nil {
		panic(err)
	} else {
		return val, nil
	}
}

func AddKeyValuePair(key string, val string, expireTime int) {
	if(useCache!="TRUE"){
		return
	}
	ctx := context.Background()
	err := RedisClient.Set(ctx, key, val, time.Duration(expireTime)*time.Second).Err()
	if err != nil {
		panic(err)
	}
}
