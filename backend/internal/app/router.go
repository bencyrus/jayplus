package app

import (
	"backend/internal/auth"

	"github.com/gorilla/mux"
)

func (app *App) SetupRouter() *mux.Router {
	r := mux.NewRouter()
	auth.AuthRoutes(r, app.Auth, app.DB)
	return r
}
