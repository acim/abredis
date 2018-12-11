package abredis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

// Client ...
type Client struct {
	*redis.Client
}

// NewClient creates new Redis client.
func NewClient(c *redis.Client) *Client {
	r := &Client{Client: c}
	r.Do("CONFIG", "SET", "notify-keyspace-events", "KEA")
	return r
}

// WatchKey returns channel with events weather provided key has been modified.
func (c *Client) WatchKey(ctx context.Context, key string) <-chan struct{} {
	pubsub := c.PSubscribe(fmt.Sprintf("__key*__:%s", key))

	res := make(chan struct{})
	go func() {
		for {
			select {
			case <-pubsub.Channel():
				res <- struct{}{}
			case <-ctx.Done():
				pubsub.Close()
			}
		}
	}()

	return res
}
