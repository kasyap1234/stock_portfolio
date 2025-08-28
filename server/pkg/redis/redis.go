package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis"
)


type RedisClient interface {
	Set(ctx context.Context,key string,value interface{},expiration time.Duration)*redis.StatusCmd
	Get(ctx context.Context,key string)*redis.StringCmd
	Del(ctx context.Context,keys ...string)*redis.IntCmd
	Exists(ctx context.Context,keys ...string)*redis.IntCmd
	Ping(ctx context.Context)*redis.StatusCmd
}

type redisClient struct {
	client *redis.Client 

}

func NewClient(addr ,password string, db int)RedisClient{
	client :=redis.NewClient(&redis.Options{
		Addr:addr,
		Password:password,
		DB:db,
	},
	)
	return &redisClient{client: client}
}

