package authentication

import (
	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, auth *Auth) {
	r.HandleFunc("/login", auth.LoginHandler).Methods("GET")
	r.HandleFunc("/authenticate", auth.Authenticate).Methods("GET")
}
