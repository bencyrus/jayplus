package app

import (
	"backend/internal/db"
)

func SetupApp() error {
	// Setup DB
	err := db.SetupDB()
	if err != nil {
		return err
	}

	// other setup stuff here

	return nil
}
