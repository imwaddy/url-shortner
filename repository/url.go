package repository

import (
	"context"
	"database/sql"

	"github.com/imwaddy/url-shortner/pkg/cache"
	"github.com/imwaddy/url-shortner/pkg/logger"
)

type URLRepository struct {
	db    *sql.DB
	cache cache.RedisClient
}

func NewURLRepository(db *sql.DB, cache cache.RedisClient) *URLRepository {
	return &URLRepository{db: db, cache: cache}
}

func (r *URLRepository) Save(short, original string) error {
	_, err := r.db.Exec(
		"INSERT INTO urls (short_code, original_url) VALUES (?, ?)",
		short, original,
	)
	return err
}

func (r *URLRepository) Get(ctx context.Context, short string) (string, error) {
	var original string

	original, err := r.cache.Get(ctx, short)
	if err == nil && original != "" {
		return original, nil
	}

	err = r.db.QueryRow(
		"SELECT original_url FROM urls WHERE short_code=?",
		short,
	).Scan(&original)
	if err != nil {
		logger.Error("Error while getting value")
		return "", err
	}

	if err = r.cache.Set(ctx, short, original); err != nil {
		logger.Errorf("Error while setting value in cache %+v", err)
	}

	return original, err
}
