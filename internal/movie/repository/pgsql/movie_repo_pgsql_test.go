package pgsql

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zakariawahyu/go-api-movie/internal/domain"
	"regexp"
	"testing"
	"time"
)

func TestMovieRepositoryPgsql_Fetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' database connection", err)
	}
	defer db.Close()

	mockMovies := []domain.Movie{
		{ID: 1, Title: "Movie 1", Description: "Description 1", Rating: 5, Image: "Image 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Title: "Movie 2", Description: "Description 2", Rating: 5, Image: "Image 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
		AddRow(mockMovies[0].ID, mockMovies[0].Title, mockMovies[0].Description, mockMovies[0].Rating, mockMovies[0].Image, mockMovies[0].CreatedAt, mockMovies[0].UpdatedAt).
		AddRow(mockMovies[1].ID, mockMovies[1].Title, mockMovies[1].Description, mockMovies[1].Rating, mockMovies[1].Image, mockMovies[1].CreatedAt, mockMovies[1].UpdatedAt)

	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM movies"
	mock.ExpectQuery(query).WillReturnRows(rows)

	movieRepo := NewMovieRepositoryPgsql(db)
	movies, err := movieRepo.Fetch(context.TODO())
	assert.NoError(t, err)
	assert.Len(t, movies, 2)
}

func TestMovieRepositoryPgsql_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' database connection", err)
	}
	defer db.Close()

	movieMock := &domain.Movie{
		ID:          1,
		Title:       "Title",
		Description: "Description",
		Rating:      5,
		Image:       "Image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
		AddRow(movieMock.ID, movieMock.Title, movieMock.Description, movieMock.Rating, movieMock.Image, movieMock.CreatedAt, movieMock.UpdatedAt)

	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM movies WHERE id = $1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(movieMock.ID).
		WillReturnRows(rows)

	movieRepo := NewMovieRepositoryPgsql(db)
	movie, err := movieRepo.GetByID(context.TODO(), movieMock.ID)
	assert.NoError(t, err)
	assert.NotNil(t, movie)
	assert.Equal(t, movieMock.ID, movie.ID)
}

func TestMovieRepositoryPgsql_CreateErr(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' database connection", err)
	}
	defer db.Close()

	movieMock := &domain.Movie{
		Title:       "Title",
		Description: "Description",
		Rating:      5,
		Image:       "Image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
		AddRow(movieMock.ID, movieMock.Title, movieMock.Description, movieMock.Rating, movieMock.Image, movieMock.CreatedAt, movieMock.UpdatedAt)

	query := "INSERT INTO movies (title, description, rating, image, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(movieMock.Title, movieMock.Description, movieMock.Rating, movieMock.Image, movieMock.CreatedAt, movieMock.UpdatedAt).
		WillReturnRows(rows)

	movieRepo := NewMovieRepositoryPgsql(db)
	_, err = movieRepo.Create(context.TODO(), movieMock)
	assert.Error(t, err)
}

func TestMovieRepositoryPgsql_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' database connection", err)
	}
	defer db.Close()

	movieMock := &domain.Movie{
		ID:          1,
		Title:       "Title",
		Description: "Description",
		Rating:      5,
		Image:       "Image",
		UpdatedAt:   time.Now(),
	}

	query := "UPDATE movies SET title = $1, description = $2, rating = $3, image = $4, updated_at = $5 WHERE id = $6"
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(movieMock.Title, movieMock.Description, movieMock.Rating, movieMock.Image, movieMock.UpdatedAt, movieMock.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	movieRepo := NewMovieRepositoryPgsql(db)
	movie, err := movieRepo.Update(context.TODO(), movieMock)
	assert.NoError(t, err)
	assert.NotNil(t, movie)
}
