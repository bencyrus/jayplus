package routes

import (
	"backend/pkg/authentication/handlers"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	// other routes here

	// dummy get route
	r.HandleFunc("/test", handlers.TestHandler).Methods("GET")
}
