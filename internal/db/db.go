package db

import (
	"api/internal/config"
	"database/sql"

	_ "github.com/lib/pq"
)

// Connect open connection with database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DBStringConnection)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
