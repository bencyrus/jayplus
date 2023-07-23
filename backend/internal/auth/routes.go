package auth

import (
	"backend/internal/db"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, auth AuthInterface, db db.DBInterface) {
	r.HandleFunc("/login", auth.LoginHandler).Methods("GET")
	r.HandleFunc("/authenticate", auth.Authenticate(db)).Methods("POST")
}
