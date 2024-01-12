package pgsql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zakariawahyu/go-api-movie/internal/domain"
)

type movieRepositoryPgsql struct {
	db *sql.DB
}

func NewMovieRepositoryPgsql(db *sql.DB) *movieRepositoryPgsql {
	return &movieRepositoryPgsql{
		db: db,
	}
}
func (r *movieRepositoryPgsql) Fetch(ctx context.Context) ([]domain.Movie, error) {
	movies := []domain.Movie{}
	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM movies"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return movies, err
	}

	defer rows.Close()

	for rows.Next() {
		movie := domain.Movie{}
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Rating, &movie.Image, &movie.CreatedAt, &movie.UpdatedAt)
		if err != nil {
			return movies, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func (r *movieRepositoryPgsql) GetByID(ctx context.Context, id int64) (*domain.Movie, error) {
	movie := &domain.Movie{}
	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM movies WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Rating, &movie.Image, &movie.CreatedAt, &movie.UpdatedAt)

	return movie, err
}
func (r *movieRepositoryPgsql) Create(ctx context.Context, movie *domain.Movie) (*domain.Movie, error) {
	query := "INSERT INTO movies (title, description, rating, image, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.ExecContext(ctx, query, movie.Title, movie.Description, movie.Rating, movie.Image, movie.CreatedAt, movie.UpdatedAt)

	return movie, err
}

func (r *movieRepositoryPgsql) Update(ctx context.Context, movie *domain.Movie) (*domain.Movie, error) {
	query := "UPDATE movies SET title = $1, description = $2, rating = $3, image = $4, updated_at = $5 WHERE id = $6"
	res, err := r.db.ExecContext(ctx, query, movie.Title, movie.Description, movie.Rating, movie.Image, movie.UpdatedAt, movie.ID)

	if err != nil {
		return nil, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected != 1 {
		return nil, fmt.Errorf("error, total affected: %d", affected)
	}

	return movie, nil
}

func (r *movieRepositoryPgsql) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM movies WHERE id = $1"
	res, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return fmt.Errorf("error, total affected: %d", affected)
	}

	return nil
}
