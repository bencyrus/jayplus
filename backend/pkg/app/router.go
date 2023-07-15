package app

import (
	"backend/pkg/authentication/routes"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Setup routes
	routes.AuthRoutes(r)

	return r
}
