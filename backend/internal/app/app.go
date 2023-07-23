package app

import (
	"backend/config"
	authDomain "backend/domains/auth"
	dbDomain "backend/domains/db"
	"backend/internal/auth"
	"backend/internal/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	DB     dbDomain.DBInterface
	Router *mux.Router
	Auth   authDomain.AuthInterface
}

func NewApp() (*App, error) {
	// Setup DB
	dbInstance := &db.DB{}
	err := dbInstance.SetupDB()
	if err != nil {
		return nil, fmt.Errorf("failed to setup the database: %w", err)
	}

	// Setup Authentication
	authInstance := auth.NewAuth()

	// Setup App
	app := &App{
		DB:     dbInstance,
		Auth:   authInstance,
		Router: mux.NewRouter(),
	}

	return app, nil
}

func (app *App) Run() error {
	fmt.Printf("Server running at %s\n", config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Port), app.Router)
}
