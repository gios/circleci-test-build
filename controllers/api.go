package controllers

import (
	"net/http"
	"os"
)

// APIVersionController - return current API version.
func APIVersionController(w http.ResponseWriter, r *http.Request) {
	apiVersion := os.Getenv("API_VERSION")
	w.Write([]byte("API Version " + apiVersion))
}