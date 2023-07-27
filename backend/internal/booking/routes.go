package booking

import (
	"net/http"

	authContracts "backend/contracts/auth"

	"github.com/gorilla/mux"
)

func BookingRoutes(r *mux.Router, authMiddleware authContracts.AuthInterface) {
	bookingRouter := r.PathPrefix("/booking").Subrouter()

	bookingRouter.Use(authMiddleware.AuthRequired)

	bookingRouter.HandleFunc("/vehicle-types", func(w http.ResponseWriter, r *http.Request) {
		GetVehicleTypes(w, r)
	}).Methods("GET")
}
