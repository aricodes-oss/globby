package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Context = context.Background()
var client *redis.Client = nil

func Client() *redis.Client {
	if client != nil {
		return client
	}

	client = redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "", // No password, docker internal network only
		DB:       0,  // Use default DB
	})

	return client
}
