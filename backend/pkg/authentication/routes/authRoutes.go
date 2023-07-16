package routes

import (
	"backend/pkg/authentication/handlers"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET")
}
