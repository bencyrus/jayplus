package auth

import (
	"backend/internal/db"
	"net/http"
)

type AuthInterface interface {
	InitAuth(auth Auth) *Auth
	GenerateSignedTokenPair(user *AuthUser) (JWTTokenPair, error)
	GetRefreshCookie(refreshToken string) *http.Cookie
	LoginHandler(w http.ResponseWriter, r *http.Request)
	Authenticate(w http.ResponseWriter, r *http.Request, db db.DBInterface)
}
