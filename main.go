package main

import (
	"context"
	"fmt"

	"github.com/acim/redis-watch/pkg/abredis"
	"github.com/go-redis/redis"
)

// https://redis.io/topics/notifications

const key = "config:app"

func main() {
	client := abredis.NewClient(redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	}))

	for _ = range client.WatchKey(context.TODO(), key) {
		fmt.Printf("key %s modified\n", key)
	}

	fmt.Println("END")
}
