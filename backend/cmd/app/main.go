package main

import (
	"backend/config"
	"backend/internal/app"
	"log"
)

func main() {
	// Load config
	config.Load()

	// Setup App
	myApp, err := app.NewApplication()
	if err != nil {
		log.Fatalf("failed to setup the application: %s", err.Error())
		return
	}

	// Run server
	err = myApp.Run()
	if err != nil {
		log.Fatalf("failed to start the server: %s", err.Error())
	}
}
