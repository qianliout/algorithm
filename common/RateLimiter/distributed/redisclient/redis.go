package redisclient

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

// NewClientFromAddr constructs a Redis client from a host:port address and optional password.
// db selects the Redis database index.
func NewClientFromAddr(addr string, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: db})
}

// Ping pings the server to ensure connectivity.
func Ping(ctx context.Context, c *redis.Client) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return c.Ping(ctx).Err()
}
