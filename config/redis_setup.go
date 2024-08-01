package config

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

var (
	client      *redis.Client
	clientOnce  sync.Once
	clientError error
)

func GetRedisClient() *redis.Client {
	clientOnce.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     REDIS_URL,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		// Ping Redis to check if the connection is working
		_, clientError := client.Ping().Result()
		if clientError != nil {
			panic(clientError)
		}
		fmt.Println("Connected to Redis")
	})
	return client
}
