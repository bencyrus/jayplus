package app

import (
	"backend/pkg/authentication"

	"github.com/gorilla/mux"
)

func (app *Application) NewRouter() *mux.Router {
	r := mux.NewRouter()
	authentication.AuthRoutes(r, app.Auth)
	return r
}
