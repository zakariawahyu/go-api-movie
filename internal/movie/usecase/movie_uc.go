package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/zakariawahyu/go-api-movie/internal/domain"
	"github.com/zakariawahyu/go-api-movie/internal/transport/request"
	"github.com/zakariawahyu/go-api-movie/utils/response"
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

	movies := []domain.Movie{}
	movieCached, _ := u.redisRepo.Get(ctx, "movies")
	if err := json.Unmarshal([]byte(movieCached), &movies); err == nil {
		return movies, err
	}

	res, err := u.movieRepo.Fetch(c)
	if err != nil {
		return nil, err
	}

	movieString, _ := json.Marshal(&res)
	if err = u.redisRepo.Set(ctx, "movies", movieString, 30*time.Second); err != nil {
		return nil, err
	}
	return res, nil
}

func (u *movieUsecase) GetByID(ctx context.Context, id int64) (*domain.Movie, error) {
	c, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	movie := &domain.Movie{}
	todosCached, _ := u.redisRepo.Get(ctx, fmt.Sprintf("movie-%d", id))
	if err := json.Unmarshal([]byte(todosCached), movie); err == nil {
		return movie, err
	}

	res, err := u.movieRepo.GetByID(c, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	movieString, _ := json.Marshal(&res)
	if err = u.redisRepo.Set(ctx, fmt.Sprintf("movie-%d", res.ID), movieString, 30*time.Second); err != nil {
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

	movieString, _ := json.Marshal(&res)
	if err = u.redisRepo.Set(ctx, fmt.Sprintf("movie-%d", res.ID), movieString, 30*time.Second); err != nil {
		return nil, err
	}

	return res, nil
}

func (u *movieUsecase) Update(ctx context.Context, req request.UpdateMovieRequest, id int64) (*domain.Movie, error) {
	c, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	movie, err := u.movieRepo.GetByID(c, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
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
		if err == sql.ErrNoRows {
			return response.ErrNotFound
		}
		return err
	}

	if err := u.movieRepo.Delete(c, id); err != nil {
		return err
	}

	return nil
}
