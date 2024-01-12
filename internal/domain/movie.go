package domain

import (
	"context"
	"github.com/zakariawahyu/go-api-movie/internal/transport/request"
	"time"
)

type Movie struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type MovieRepositoryPgsql interface {
	Fetch(ctx context.Context) ([]Movie, error)
	GetByID(ctx context.Context, id int64) (*Movie, error)
	Create(ctx context.Context, req *Movie) (*Movie, error)
	Update(ctx context.Context, req *Movie) (*Movie, error)
	Delete(ctx context.Context, id int64) error
}

type MovieRepositoryRedis interface {
	Delete(ctx context.Context, key string) error
	Set(ctx context.Context, key string, value interface{}, exp time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

type MovieUsecase interface {
	Fetch(ctx context.Context) ([]Movie, error)
	GetByID(ctx context.Context, id int64) (*Movie, error)
	Create(ctx context.Context, req request.CreateMovieRequest) (*Movie, error)
	Update(ctx context.Context, req request.UpdateMovieRequest, id int64) (*Movie, error)
	Delete(ctx context.Context, id int64) error
}
