package handlers

import (
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	// dummy handler
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("test handler"))
}
