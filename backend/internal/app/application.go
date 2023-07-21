package app

import (
	"backend/config"
	"backend/internal/authentication"
	"backend/internal/db"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Application struct {
	DB     *sql.DB
	Router *mux.Router
	Auth   *authentication.Auth
}

func NewApplication() (*Application, error) {
	// Setup DB
	db, err := db.SetupDB()
	if err != nil {
		return nil, err
	}

	// Setup Application
	app := &Application{
		DB: db,
	}

	// Setup Authentication
	app.Auth = authentication.NewAuth(authentication.Auth{
		Issuer:             config.JWTIssuer,
		Audience:           config.JWTAudience,
		Secret:             config.JWTSecret,
		AccessTokenExpiry:  config.AccessTokenExpiry,
		RefreshTokenExpiry: config.RefreshTokenExpiry,
		CookieDomain:       config.JWTCookieDomain,
		CookiePath:         config.JWTCookiePath,
		CookieName:         config.JWTCookieName,
	})

	// Setup Router
	app.Router = app.NewRouter()

	// Return new Application
	return app, nil
}

func (app *Application) Run() error {
	fmt.Printf("Server running at %s\n", config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Port), app.Router)
}
