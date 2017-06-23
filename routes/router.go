package router

import (
	"net/http"

	ctrl "github.com/gios/mom-and-i/controllers"
	"github.com/gorilla/mux"
)

// InitRoutes - initialize routes for app
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/products").Methods("GET").HandlerFunc(ctrl.ProductsController)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	return router
}
