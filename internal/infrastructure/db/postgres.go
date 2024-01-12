package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/zakariawahyu/go-api-movie/config"
	"os"
	"time"
)

func NewPostgresConnection(cfg *config.Config) (*sql.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName)
	db, err := sql.Open("postgres", url)
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(5)

	return db, nil
}
