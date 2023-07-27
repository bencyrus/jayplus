package auth

import (
	dbContracts "backend/contracts/db"
	"net/http"
)

type AuthHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request, db dbContracts.DBInterface)
	RefreshToken(w http.ResponseWriter, r *http.Request, db dbContracts.DBInterface)
	Logout(w http.ResponseWriter, r *http.Request)
}

type AuthMiddlewareInterface interface {
	AuthRequired(next http.Handler) http.Handler
}
