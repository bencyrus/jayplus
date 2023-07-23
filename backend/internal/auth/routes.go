package auth

import (
	authDomain "backend/domains/auth"
	dbDomain "backend/domains/db"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, auth authDomain.AuthInterface, db dbDomain.DBInterface) {
	r.HandleFunc("/login", auth.LoginHandler).Methods("GET")
	r.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		auth.Authenticate(w, r, db)
	})
}
