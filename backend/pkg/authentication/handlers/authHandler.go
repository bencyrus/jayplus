package handlers

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// dummy handler
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("login handler"))
}
