package internal

import (
	"database/sql"
)

func OpenDB(connectionString string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
