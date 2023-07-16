package tests

import (
	"backend/config"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// test cases
	cases := []struct {
		Name               string
		EnvVars            map[string]string
		ExpectedDBHost     string
		ExpectedDBPort     string
		ExpectedDBName     string
		ExpectedDBUser     string
		ExpectedDBPassword string
		ExpectedPort       string
	}{
		{
			Name: "All variables set",
			EnvVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_NAME":     "test_db",
				"DB_USER":     "test_user",
				"DB_PASSWORD": "test_password",
				"PORT":        "8080",
			},
			ExpectedDBHost:     "localhost",
			ExpectedDBPort:     "5432",
			ExpectedDBName:     "test_db",
			ExpectedDBUser:     "test_user",
			ExpectedDBPassword: "test_password",
			ExpectedPort:       "8080",
		},
		{
			Name:               "No variables set",
			EnvVars:            map[string]string{},
			ExpectedDBHost:     "",
			ExpectedDBPort:     "",
			ExpectedDBName:     "",
			ExpectedDBUser:     "",
			ExpectedDBPassword: "",
			ExpectedPort:       "",
		},
		{
			Name: "Some variables set",
			EnvVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_NAME":     "test_db",
				"DB_PASSWORD": "test_password",
			},
			ExpectedDBHost:     "localhost",
			ExpectedDBPort:     "",
			ExpectedDBName:     "test_db",
			ExpectedDBUser:     "",
			ExpectedDBPassword: "test_password",
			ExpectedPort:       "",
		},
		{
			Name: "Only DB related variables set",
			EnvVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_NAME":     "test_db",
				"DB_USER":     "test_user",
				"DB_PASSWORD": "test_password",
			},
			ExpectedDBHost:     "localhost",
			ExpectedDBPort:     "5432",
			ExpectedDBName:     "test_db",
			ExpectedDBUser:     "test_user",
			ExpectedDBPassword: "test_password",
			ExpectedPort:       "",
		},
	}

	// Iterate over test cases
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			// Set environment variables
			for k, v := range tc.EnvVars {
				os.Setenv(k, v)
			}

			// Load config
			config.Load()

			// Check config values
			if config.DBHost != tc.ExpectedDBHost ||
				config.DBPort != tc.ExpectedDBPort ||
				config.DBName != tc.ExpectedDBName ||
				config.DBUser != tc.ExpectedDBUser ||
				config.DBPassword != tc.ExpectedDBPassword ||
				config.Port != tc.ExpectedPort {
				t.Fatalf("Loaded config does not match expected config")
			}

			// Unset environment variables for next run
			for k := range tc.EnvVars {
				os.Unsetenv(k)
			}
		})
	}
}
