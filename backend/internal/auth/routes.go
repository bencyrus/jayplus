package auth

import (
	"backend/internal/db"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, auth AuthInterface, db db.DBInterface) {
	r.HandleFunc("/login", auth.LoginHandler).Methods("GET")
	r.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		auth.Authenticate(w, r, db)
	}).Methods("POST")
}
