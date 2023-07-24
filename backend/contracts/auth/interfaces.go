package auth

import (
	dbContracts "backend/contracts/db"
	"net/http"
)

type AuthInterface interface {
	LoginHandler(w http.ResponseWriter, r *http.Request)
	Authenticate(w http.ResponseWriter, r *http.Request, db dbContracts.DBInterface)
}
