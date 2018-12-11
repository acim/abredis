package main

import (
	"context"
	"fmt"

	"github.com/acim/redis-watch/pkg/abredis"
	"github.com/go-redis/redis"
)

const key = "config:app"

func main() {
	client := abredis.NewClient(redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	}))

	for _ = range client.KeyW(context.TODO(), key) {
		fmt.Printf("key %s modified\n", key)
	}
}
