package redis

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/zakariawahyu/go-api-movie/internal/domain"
	"log"
	"testing"
	"time"
)

func SetupRedis() domain.MovieRepositoryRedis {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub redis connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	redisRepository := NewMovieRepositoryRedis(client)
	return redisRepository
}

func TestSet(t *testing.T) {
	redisRepository := SetupRedis()
	err := redisRepository.Set(context.TODO(), "ping", 0, time.Duration(0))
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	redisRepository := SetupRedis()
	key, val, exp := "ping", "pong", time.Duration(0)

	value, err := redisRepository.Get(context.TODO(), key)
	assert.NotNil(t, err)
	assert.Equal(t, value, "")

	err = redisRepository.Set(context.TODO(), key, val, exp)
	assert.NoError(t, err)

	value, err = redisRepository.Get(context.TODO(), key)
	assert.NoError(t, err)
	assert.Equal(t, value, val)
}
