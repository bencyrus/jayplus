package app

import (
	"backend/config"
	"backend/internal/db"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Application struct {
	DB     *sql.DB
	Router *mux.Router
}

func NewApplication() (*Application, error) {
	// Setup DB
	db, err := db.SetupDB()
	if err != nil {
		return nil, err
	}

	// Setup Router
	r := NewRouter()

	// Return new Application
	return &Application{
		DB:     db,
		Router: r,
	}, nil
}

func (app *Application) Run() error {
	fmt.Printf("Server running at %s\n", config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Port), app.Router)
}
