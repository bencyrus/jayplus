package db

import "backend/models"

type DBInterface interface {
	SetupDB() error
	GetUserByEmail(email string) (*models.User, error)
}
