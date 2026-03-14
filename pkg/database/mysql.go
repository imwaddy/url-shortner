package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imwaddy/url-shortner/pkg/logger"
)

func NewMySQL(user, pass, host, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		user, pass, host, dbName,
	)

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}

		logger.Println("Waiting for MySQL...")
		time.Sleep(3 * time.Second)
	}

	return nil, err
}

func AutoMigrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id INT AUTO_INCREMENT PRIMARY KEY,
		short_code VARCHAR(10) NOT NULL UNIQUE,
		original_url TEXT NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		logger.Errorf("Error while automigrate %+v", err)
		return err
	}
	return err
}
