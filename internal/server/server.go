package server

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-api-movie/config"
	"github.com/zakariawahyu/go-api-movie/internal/infrastructure/cache"
	"github.com/zakariawahyu/go-api-movie/internal/infrastructure/db"
	v1 "github.com/zakariawahyu/go-api-movie/internal/movie/delivery/http/v1"
	"github.com/zakariawahyu/go-api-movie/internal/movie/repository/pgsql"
	redisRepo "github.com/zakariawahyu/go-api-movie/internal/movie/repository/redis"
	"github.com/zakariawahyu/go-api-movie/internal/movie/usecase"
	"time"
)

type server struct {
	cfg  *config.Config
	echo *echo.Echo
}

func NewHttpServer(cfg *config.Config) *server {
	return &server{
		cfg:  cfg,
		echo: echo.New(),
	}
}

func (s *server) Run() error {
	db, err := db.NewPostgresConnection(s.cfg)
	if err != nil {
		return err
	}

	redis := cache.NewRedisConnection(s.cfg)

	movieRepo := pgsql.NewMovieRepositoryPgsql(db)
	redisRepo := redisRepo.NewMovieRepositoryRedis(redis)

	ctxTimeout := time.Duration(s.cfg.Server.ReadTimeout) * time.Second
	movieUsecase := usecase.NewMovieUsecase(movieRepo, redisRepo, ctxTimeout)

	v1.NewMovieHandler(s.echo, movieUsecase)

	if err := s.runHttpServer(); err != nil {
		s.echo.Logger.Errorf(" err run httpserver: %v", err)
	}

	return nil
}
