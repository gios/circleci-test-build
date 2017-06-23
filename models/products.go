package models

import (
	"database/sql"

	dbcon "github.com/gios/mom-and-i/shared/database"
	log "github.com/sirupsen/logrus"
)

// GetProducts - gets all data about products
func GetProducts() *sql.Rows {
	db, err := dbcon.Connect()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT name FROM products")
	if err != nil {
		log.Error(err)
	}
	return rows
}
