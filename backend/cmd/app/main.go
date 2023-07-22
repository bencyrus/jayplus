package main

import (
	"backend/config"
	"backend/internal/app"
	"backend/internal/auth"
	"backend/internal/db"
	"log"
)

func main() {
	// Load config
	config.Load()

	// Setup DB
	dbInstance := &db.DB{}
	err := dbInstance.SetupDB()
	if err != nil {
		log.Fatalf("failed to setup the database: %s", err.Error())
		return
	}

	// Setup Authentication
	authInstance := &auth.Auth{}
	authInstance = authInstance.InitAuth(auth.Auth{
		Issuer:             config.JWTIssuer,
		Audience:           config.JWTAudience,
		Secret:             config.JWTSecret,
		AccessTokenExpiry:  config.AccessTokenExpiry,
		RefreshTokenExpiry: config.RefreshTokenExpiry,
		CookieDomain:       config.JWTCookieDomain,
		CookiePath:         config.JWTCookiePath,
		CookieName:         config.JWTCookieName,
	})

	// Setup App
	appInstance, err := app.InitApp(dbInstance, authInstance)
	if err != nil {
		log.Fatalf("failed to setup the application: %s", err.Error())
		return
	}

	// Run server
	err = appInstance.Run()
	if err != nil {
		log.Fatalf("failed to start the server: %s", err.Error())
	}
}
