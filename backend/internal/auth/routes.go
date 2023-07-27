package auth

import (
	"net/http"

	authContracts "backend/contracts/auth"
	dbContracts "backend/contracts/db"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, auth authContracts.AuthHandlerInterface, db dbContracts.DBInterface) {
	authRouter := r.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", auth.LoginHandler).Methods("GET")
	authRouter.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		auth.Authenticate(w, r, db)
	}).Methods("POST")
	authRouter.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		auth.RefreshToken(w, r, db)
	}).Methods("GET")
	authRouter.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		auth.Logout(w, r)
	}).Methods("GET")
}
