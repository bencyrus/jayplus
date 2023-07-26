package auth

import (
	"net/http"

	authContracts "backend/contracts/auth"
	dbContracts "backend/contracts/db"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, auth authContracts.AuthInterface, db dbContracts.DBInterface) {
	r.HandleFunc("/login", auth.LoginHandler).Methods("GET")
	r.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		auth.Authenticate(w, r, db)
	}).Methods("POST")
	r.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		auth.RefreshToken(w, r, db)
	}).Methods("GET")
	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		auth.Logout(w, r)
	}).Methods("GET")
}
