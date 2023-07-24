package db

import (
	"backend/config"
	"backend/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	dbConnectTimeout = 3
)

type DB struct {
	*sql.DB
}

func (db *DB) SetupDB() error {
	pgConString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)

	var err error
	db.DB, err = sql.Open("pgx", pgConString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	log.Println("Database connection successfully established")

	return nil
}

// Get user by email
func (db *DB) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbConnectTimeout*time.Second)
	defer cancel()

	quey := `SELECT 
				id,
				email,
				phone_number,
				password, 
				first_name,
				last_name,
				role,
				created_at,
				updated_at
			FROM 
				users
			WHERE
				email = $1`

	var user models.User

	row := db.QueryRowContext(ctx, quey, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PhoneNumber,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Get user by ID
func (db *DB) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbConnectTimeout*time.Second)
	defer cancel()

	quey := `SELECT 
				id,
				email,
				phone_number,
				password, 
				first_name,
				last_name,
				role,
				created_at,
				updated_at
			FROM
				users
			WHERE
				id = $1`

	var user models.User

	row := db.QueryRowContext(ctx, quey, id)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PhoneNumber,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
