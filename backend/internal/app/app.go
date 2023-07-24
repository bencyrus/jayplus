package app

import (
	"backend/config"
	"backend/internal/auth"
	"backend/internal/db"
	"fmt"
	"net/http"

	authContracts "backend/contracts/auth"
	dbContracts "backend/contracts/db"

	"github.com/gorilla/mux"
)

type App struct {
	DB     dbContracts.DBInterface
	Router *mux.Router
	Auth   authContracts.AuthInterface
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
		DB:   dbInstance,
		Auth: authInstance,
	}
	app.Router = app.SetupRouter()

	return app, nil
}

func (app *App) Run() error {
	fmt.Printf("Server running at %s\n", config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Port), app.Router)
}
