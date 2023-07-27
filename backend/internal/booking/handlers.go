package booking

import (
	"backend/utils"
	"net/http"

	bookingContracts "backend/contracts/booking"
)

// a handler that returns a list of vehicle types. the tpyes are Sedan, SUV, Large SUV/Truck, and Motorcycle
func GetVehicleTypes(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, []bookingContracts.VehicleType{
		{ID: 1, Name: "Sedan"},
		{ID: 2, Name: "SUV"},
		{ID: 3, Name: "Large SUV/Truck"},
		{ID: 4, Name: "Motorcycle"},
	})
}
