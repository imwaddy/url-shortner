package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/imwaddy/url-shortner/pkg/logger"
)

type RedisClientStruct struct {
	cache *redis.Client
}

type RedisClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
	Close() error
}

func NewRedisClient(addr string) RedisClient {
	redisClient := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	for i := 0; i < 10; i++ {
		err := redisClient.Ping(context.Background()).Err()
		if err == nil {
			logger.Println("✅ Connected to Redis")
			break
		}

		logger.Println("⏳ Waiting for Redis...")
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

func (r *RedisClientStruct) Close() error {
	return r.cache.Close()
}
