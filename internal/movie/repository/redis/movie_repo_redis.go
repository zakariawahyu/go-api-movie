package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-api-movie/internal/domain"
	"time"
)

type movieRepositoryRedis struct {
	redis *redis.Client
}

func NewMovieRepositoryRedis(redis *redis.Client) domain.MovieRepositoryRedis {
	return &movieRepositoryRedis{
		redis: redis,
	}
}

func (r *movieRepositoryRedis) Delete(ctx context.Context, key string) error {
	if err := r.redis.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}

func (r *movieRepositoryRedis) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	return r.redis.Set(ctx, key, value, exp).Err()
}

func (r *movieRepositoryRedis) Get(ctx context.Context, key string) (string, error) {
	return r.redis.Get(ctx, key).Result()
}
