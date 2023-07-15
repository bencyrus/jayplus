package main

import (
	"backend/pkg/app"
	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalf("failed to start the application: %s", err.Error())
	}
}
