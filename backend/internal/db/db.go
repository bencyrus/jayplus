package db

import (
	"backend/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func SetupDB() (*sql.DB, error) {
	pgConString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)

	db, err := sql.Open("pgx", pgConString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connection successfully established")

	return db, nil
}
