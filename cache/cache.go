package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Context = context.Background()
var Client *redis.Client = nil

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "", // No password, docker internal network only
		DB:       0,  // Use default DB
	})
}
