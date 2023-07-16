package tests

import (
	"backend/config"
	"backend/internal/db"
	"os"
	"testing"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var dbConfig DBConfig

func init() {
	dbConfig = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func TestSetupDB(t *testing.T) {
	// Define test cases
	cases := []struct {
		Name     string
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		WantErr  bool
	}{
		{
			Name:     "Successful connection",
			Host:     dbConfig.Host,
			Port:     dbConfig.Port,
			User:     dbConfig.User,
			Password: dbConfig.Password,
			DBName:   dbConfig.DBName,
			WantErr:  false,
		},
		{
			Name:     "Failed connection due to wrong host",
			Host:     "wrong_host",
			Port:     dbConfig.Port,
			User:     dbConfig.User,
			Password: dbConfig.Password,
			DBName:   dbConfig.DBName,
			WantErr:  true,
		},
		{
			Name:     "Failed connection due to wrong port",
			Host:     dbConfig.Host,
			Port:     "wrong_port",
			User:     dbConfig.User,
			Password: dbConfig.Password,
			DBName:   dbConfig.DBName,
			WantErr:  true,
		},
		{
			Name:     "Failed connection due to wrong database name",
			Host:     dbConfig.Host,
			Port:     dbConfig.Port,
			User:     dbConfig.User,
			Password: dbConfig.Password,
			DBName:   "wrong_db_name",
			WantErr:  true,
		},
		{
			Name:     "Failed connection due to wrong username",
			Host:     dbConfig.Host,
			Port:     dbConfig.Port,
			User:     "wrong_user",
			Password: dbConfig.Password,
			DBName:   dbConfig.DBName,
			WantErr:  true,
		},
		{
			Name:     "Failed connection due to wrong password",
			Host:     dbConfig.Host,
			Port:     dbConfig.Port,
			User:     dbConfig.User,
			Password: "wrong_password",
			DBName:   dbConfig.DBName,
			WantErr:  true,
		},
		// Add more cases if needed
	}

	// Iterate over test cases
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			// Set environment variables
			os.Setenv("DB_HOST", tc.Host)
			os.Setenv("DB_PORT", tc.Port)
			os.Setenv("DB_USER", tc.User)
			os.Setenv("DB_PASSWORD", tc.Password)
			os.Setenv("DB_NAME", tc.DBName)

			// Load config
			config.Load()

			// Call SetupDB
			_, err := db.SetupDB()

			// If an error is expected
			if tc.WantErr {
				if err == nil {
					t.Fatalf("Expected error, got nil")
				}
				return
			}

			// If no error is expected
			if err != nil {
				t.Fatalf("Expected no error, got: %v", err)
			}
		})
	}
}
