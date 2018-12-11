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

// KeyW returns channel with events when the key has been modified.
func (c *Client) KeyW(ctx context.Context, key string) <-chan struct{} {
	client := c.WithContext(ctx)
	pubsub := client.PSubscribe(fmt.Sprintf("__key*__:%s", key))

	res := make(chan struct{})
	go func() {
		for {
			select {
			case <-pubsub.Channel():
				res <- struct{}{}
			case <-ctx.Done():
				pubsub.Close()
				close(res)
				return
			}
		}
	}()

	return res
}

// GetW continuously sends provided key's value after each modification.
func (c *Client) GetW(ctx context.Context, key string, data interface{}) (<-chan interface{}, <-chan error) {
	client := c.WithContext(ctx)
	pubsub := client.PSubscribe(fmt.Sprintf("__key*__:%s", key))

	res := make(chan interface{})
	errc := make(chan error)
	go func() {
		for {
			select {
			case <-pubsub.Channel():
				err := client.Get(key).Scan(data)
				if err != nil && err != redis.Nil {
					errc <- err
					continue
				}
				res <- data
			case <-ctx.Done():
				pubsub.Close()
				close(res)
				return
			}
		}
	}()

	return res, errc
}
