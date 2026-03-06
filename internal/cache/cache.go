package cache

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/imwaddy/url-shortner/internal/logger"
)

type RedisClientStruct struct {
	cache *redis.Client
}

type RedisClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

func NewRedisClient() RedisClient {
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"), // e.g. "localhost:6379"
	})

	for i := 0; i < 10; i++ {
		err := redisClient.Ping(context.Background()).Err()
		if err == nil {
			break
		}

		logger.Println("waiting for redis...")
		time.Sleep(2 * time.Second)
	}

	return &RedisClientStruct{
		cache: redisClient,
	}
}

func (r *RedisClientStruct) Get(ctx context.Context, key string) (string, error) {
	return r.cache.Get(ctx, key).Result()
}

func (r *RedisClientStruct) Set(ctx context.Context, key string, value string) error {
	// TTL 24h
	return r.cache.Set(ctx, key, value, 24*time.Hour).Err()
}
