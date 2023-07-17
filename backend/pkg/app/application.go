package app

import (
	"backend/config"
	"backend/internal/db"
	"backend/pkg/authentication"
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

	// Setup Router
	r := NewRouter()

	// Setup Authentication
	auth := authentication.SetupAuth()

	// Return new Application
	return &Application{
		DB:     db,
		Router: r,
		Auth:   auth,
	}, nil
}

func (app *Application) Run() error {
	fmt.Printf("Server running at %s\n", config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Port), app.Router)
}
