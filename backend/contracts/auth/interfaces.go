package auth

import (
	dbContracts "backend/contracts/db"
	"net/http"
)

type AuthInterface interface {
	// Handlers
	Login(w http.ResponseWriter, r *http.Request, db dbContracts.DBInterface)
	RefreshToken(w http.ResponseWriter, r *http.Request, db dbContracts.DBInterface)
	Logout(w http.ResponseWriter, r *http.Request)

	// Middleware
	AuthRequired(next http.Handler) http.Handler
}
