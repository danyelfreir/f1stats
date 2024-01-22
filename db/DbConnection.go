package db

import (
	"database/sql"
	"fmt"
	"os"
)

func OpenDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_DTBS"),
	)
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
