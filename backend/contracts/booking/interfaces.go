package booking

import "net/http"

type BookingHandlerInterface interface {
	GetVehicleTypes(w http.ResponseWriter, r *http.Request)
}
