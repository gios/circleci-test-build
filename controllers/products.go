package controllers

import (
	"net/http"

	dbModel "github.com/gios/mom-and-i/models"
)

// ProductsController - basic controller for products
func ProductsController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Products!\n"))
	dbModel.GetProducts()
}
