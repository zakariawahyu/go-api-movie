package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/zakariawahyu/go-api-movie/internal/domain"
	"github.com/zakariawahyu/go-api-movie/internal/transport/request"
	"time"
)

type movieUsecase struct {
	movieRepo domain.MovieRepositoryPgsql
	redisRepo domain.MovieRepositoryRedis
	timeout   time.Duration
}

func NewMovieUsecase(movieRepo domain.MovieRepositoryPgsql, redisRepo domain.MovieRepositoryRedis, timeout time.Duration) *movieUsecase {
	return &movieUsecase{
		movieRepo: movieRepo,
		redisRepo: redisRepo,
		timeout:   timeout,
	}
}

func (u *movieUsecase) Fetch(ctx context.Context) ([]domain.Movie, error) {
	c, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	moviesRedis, err := u.redisRepo.GetFetch(c, "movies")
	if moviesRedis != nil {
		return moviesRedis, nil
	}

	res, err := u.movieRepo.Fetch(c)
	if err != nil {
		return nil, err
	}

	err = u.redisRepo.SetFetch(c, "movies", 30, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *movieUsecase) GetByID(ctx context.Context, id int64) (*domain.Movie, error) {
	c, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	movieRedis, err := u.redisRepo.Get(c, fmt.Sprintf("movie-%d", id))
	if movieRedis != nil {
		return movieRedis, nil
	}

	res, err := u.movieRepo.GetByID(c, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("movie not found")
		}
		return nil, err
	}

	err = u.redisRepo.Set(c, fmt.Sprintf("movie-%d", id), 30, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *movieUsecase) Create(ctx context.Context, req request.CreateMovieRequest) (*domain.Movie, error) {
	c, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	movie := &domain.Movie{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	res, err := u.movieRepo.Create(c, movie)
	if err != nil {
		return nil, err
	}

	err = u.redisRepo.Set(c, fmt.Sprintf("movie-%d", res.ID), 30, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *movieUsecase) Update(ctx context.Context, req request.UpdateMovieRequest, id int64) (*domain.Movie, error) {
	c, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	movie, err := u.movieRepo.GetByID(c, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("movie not found")
		}
		return nil, err
	}

	movie.Title = req.Title
	movie.Description = req.Description
	movie.Rating = req.Rating
	movie.Image = req.Image
	movie.UpdatedAt = time.Now()

	res, err := u.movieRepo.Update(c, movie)
	if err != nil {
		return nil, err
	}

	err = u.redisRepo.Delete(c, fmt.Sprintf("movie-%d", res.ID))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *movieUsecase) Delete(ctx context.Context, id int64) error {
	c, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	_, err := u.movieRepo.GetByID(c, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("movie not found")
		}
		return err
	}

	if err := u.movieRepo.Delete(c, id); err != nil {
		return err
	}

	return nil
}
