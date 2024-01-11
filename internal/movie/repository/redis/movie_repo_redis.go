package redis

import (
	"context"
	"encoding/json"
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

func (r *movieRepositoryRedis) Set(ctx context.Context, key string, ttl int, movie *domain.Movie) error {
	movieBytes, err := json.Marshal(&movie)
	if err != nil {
		return err
	}

	if err = r.redis.Set(ctx, key, movieBytes, time.Second*time.Duration(ttl)).Err(); err != nil {
		return err
	}

	return nil
}

func (r *movieRepositoryRedis) Get(ctx context.Context, key string) (*domain.Movie, error) {
	movieBytes, err := r.redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	movie := &domain.Movie{}
	if err = json.Unmarshal(movieBytes, movie); err != nil {
		return nil, err
	}

	return movie, nil
}

func (r *movieRepositoryRedis) SetFetch(ctx context.Context, key string, ttl int, movies []domain.Movie) error {
	movieBytes, err := json.Marshal(&movies)
	if err != nil {
		return err
	}

	if err = r.redis.Set(ctx, key, movieBytes, time.Second*time.Duration(ttl)).Err(); err != nil {
		return err
	}

	return nil
}

func (r *movieRepositoryRedis) GetFetch(ctx context.Context, key string) ([]domain.Movie, error) {
	movieBytes, err := r.redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	movie := []domain.Movie{}
	if err = json.Unmarshal(movieBytes, &movie); err != nil {
		return nil, err
	}

	return movie, nil
}
