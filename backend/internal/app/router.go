package app

import (
	"backend/internal/auth"
	"backend/internal/booking"

	"github.com/gorilla/mux"
)

func (app *App) SetupRouter() *mux.Router {
	r := mux.NewRouter()
	auth.AuthRoutes(r, app.Auth, app.DB)
	booking.BookingRoutes(r)
	return r
}
