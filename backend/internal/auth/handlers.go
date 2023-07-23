package auth

import (
	"backend/internal/db"
	"backend/utils"
	"errors"
	"log"
	"net/http"
)

func (a *Auth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// dummy handler
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("login handler"))
}

func (a *Auth) Authenticate(db db.DBInterface) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqPayload struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err := utils.ReadJSON(w, r, &reqPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		user, err := db.GetUserByEmail(reqPayload.Email)
		if err != nil {
			utils.ErrorJSON(w, errors.New("invalid login credentials"), http.StatusBadRequest)
			return
		}

		valid, err := PasswordMatches(user, reqPayload.Password)
		if err != nil || !valid {
			utils.ErrorJSON(w, errors.New("invalid login credentials"), http.StatusBadRequest)
			return
		}

		// create JWT user
		u := AuthUser{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}

		// generate JWT tokens
		tokenPair, err := a.GenerateSignedTokenPair(&u)
		if err != nil {
			log.Fatalf("Error generating token pair: %v", err)
		}

		// set refresh cookie
		refreshCookie := a.GetRefreshCookie(tokenPair.RefreshToken)
		http.SetCookie(w, refreshCookie)

		utils.WriteJSON(w, http.StatusAccepted, tokenPair)
	}
}
