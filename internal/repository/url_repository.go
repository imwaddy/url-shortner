package repository

import "database/sql"

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Save(short, original string) error {
	_, err := r.db.Exec(
		"INSERT INTO urls (short_code, original_url) VALUES (?, ?)",
		short, original,
	)
	return err
}

func (r *URLRepository) Get(short string) (string, error) {
	var original string
	err := r.db.QueryRow(
		"SELECT original_url FROM urls WHERE short_code=?",
		short,
	).Scan(&original)

	return original, err
}
