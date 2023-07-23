package auth

import (
	"backend/domains/db"
	"net/http"
)

type AuthInterface interface {
	GenerateSignedTokenPair(user *AuthUser) (JWTTokenPair, error)
	GetRefreshCookie(refreshToken string) *http.Cookie
	GetExpiredRefreshCookie() *http.Cookie
	LoginHandler(w http.ResponseWriter, r *http.Request)
	Authenticate(w http.ResponseWriter, r *http.Request, db db.DBInterface)
}
