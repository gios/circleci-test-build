package database

import (
	"database/sql"
	"os"
)

// Connect - open connection to postgres
func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	return db, err
}
