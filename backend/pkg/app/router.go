package app

import (
	"backend/pkg/authentication/routes"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	routes.AuthRoutes(r)
	return r
}
