package auth

import (
	authDomain "backend/domains/auth"
	dbDomain "backend/domains/db"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, auth authDomain.AuthInterface, db dbDomain.DBInterface) {
	r.HandleFunc("/login", auth.LoginHandler).Methods("GET")
	r.HandleFunc("/authenticate", auth.Authenticate(db)).Methods("POST")
}
