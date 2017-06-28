package models

import (
	"database/sql"

	dbConnector "github.com/gios/mom-and-i/shared/database"
	log "github.com/sirupsen/logrus"
)

// GetProducts - gets all data about products
func GetProducts() *sql.Rows {
	db, err := dbConnector.Connect()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT name FROM products")
	if err != nil {
		log.Error(err)
	}
	return rows
}
