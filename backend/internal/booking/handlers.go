package booking

import (
	"backend/utils"
	"net/http"

	bookingContracts "backend/contracts/booking"
)

func GetVehicleTypes(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, []bookingContracts.VehicleType{
		{ID: 1, Name: "Sedan"},
		{ID: 2, Name: "SUV"},
		{ID: 3, Name: "Large SUV/Truck"},
		{ID: 4, Name: "Motorcycle"},
	})
}
