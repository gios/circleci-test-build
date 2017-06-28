package router

import (
	"net/http"

	ctrl "github.com/gios/mom-and-i/controllers"
	"github.com/gorilla/mux"
)

// InitRoutes - initialize routes for app
func InitRoutes() *mux.Router {
	mainRouter := mux.NewRouter()
	mainRouter.PathPrefix("/products").Methods("GET").HandlerFunc(ctrl.ProductsController)
	initApiRoutes(mainRouter)
	mainRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	return mainRouter
}

func initApiRoutes(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/version", ctrl.APIVersionController).Methods("GET")
}
