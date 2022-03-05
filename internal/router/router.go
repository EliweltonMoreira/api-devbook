package router

import (
	"api/internal/router/routes"

	"github.com/gorilla/mux"
)

// Generate return a router with configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
