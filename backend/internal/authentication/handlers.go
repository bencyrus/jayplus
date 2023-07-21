package authentication

import (
	"log"
	"net/http"
)

func (a *Auth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// dummy handler
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("login handler"))
}

func (a *Auth) Authenticate(w http.ResponseWriter, r *http.Request) {
	// create JWT user
	user := AuthUser{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}

	// generate JWT tokens
	tokenPair, err := a.GenerateSignedTokenPair(&user)
	if err != nil {
		log.Fatalf("Error generating token pair: %v", err)
	}

	// set refresh cookie
	refreshCookie := a.GetRefreshCookie(tokenPair.RefreshToken)
	http.SetCookie(w, refreshCookie)

	w.Write([]byte(tokenPair.AccessToken))
}
