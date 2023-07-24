package auth

import (
	"backend/domains/db"
	"net/http"
)

type AuthInterface interface {
	LoginHandler(w http.ResponseWriter, r *http.Request)
	Authenticate(w http.ResponseWriter, r *http.Request, db db.DBInterface)
}
