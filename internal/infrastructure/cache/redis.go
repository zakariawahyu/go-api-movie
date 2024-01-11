package cache

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-api-movie/config"
)

func NewRedisConnection(cfg *config.Config) *redis.Client {
	dsn := fmt.Sprintf("%s:%s",
		cfg.Redis.Host,
		cfg.Redis.Port,
	)

	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: cfg.Redis.Password,
	})

	return client
}
