package booking

import (
	"net/http"

	"github.com/gorilla/mux"
)

func BookingRoutes(r *mux.Router) {
	// route to vehicle types
	bookingRouter := r.PathPrefix("/booking").Subrouter()

	bookingRouter.HandleFunc("/vehicle-types", func(w http.ResponseWriter, r *http.Request) {
		GetVehicleTypes(w, r)
	}).Methods("GET")
}
