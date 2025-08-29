package rediscache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Exists(ctx context.Context, keys ...string) *redis.IntCmd
	Ping(ctx context.Context) *redis.StatusCmd
}

type redisClient struct {
	client *redis.Client
}

// NewClient creates a new Redis client wrapper
func NewClient(addr, password string, db int) RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &redisClient{client: client}
}

func (c *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.client.Set(ctx, key, value, expiration)
}

func (c *redisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return c.client.Get(ctx, key)
}

func (c *redisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return c.client.Del(ctx, keys...)
}

func (c *redisClient) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	return c.client.Exists(ctx, keys...)
}

func (c *redisClient) Ping(ctx context.Context) *redis.StatusCmd {
	return c.client.Ping(ctx)
}
