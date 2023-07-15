package app

import (
	"backend/config"
	"fmt"
	"net/http"
)

func Run() error {
	r := SetupRouter()
	fmt.Printf("Server running at %s\n", config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r)
}
