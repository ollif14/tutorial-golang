package config

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var Ctx = context.TODO()
var Client = &redisClient{}

type redisClient struct {
	 C *redis.Client
}

func RedisClient() (*redisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // host:port of the redis server
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	Client.C = client
	fmt.Println("Connected to Redis")

	return Client, nil
}

func (client *redisClient) GetKey(key string, src interface{}) error {
	val, err := client.C.Get(key).Result()
	if err == redis.Nil || err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		return err
	}
	return nil
}

func (client *redisClient) SetKey(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = client.C.Set(key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}



