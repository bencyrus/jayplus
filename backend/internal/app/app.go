package app

import (
	"backend/config"
	"backend/internal/auth"
	"backend/internal/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	DB     db.DBInterface
	Router *mux.Router
	Auth   auth.AuthInterface
}

func InitApp(db db.DBInterface, auth auth.AuthInterface) (*App, error) {
	app := &App{
		DB:   db,
		Auth: auth,
	}

	app.Router = app.InitRouter()

	return app, nil
}

func (app *App) Run() error {
	fmt.Printf("Server running at %s\n", config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Port), app.Router)
}
