package tests

import (
	"backend/config"
	"os"
	"testing"
	"time"
)

func TestLoadDBConfig(t *testing.T) {
	// test cases
	cases := []struct {
		Name               string
		EnvVars            map[string]string
		ExpectedDBHost     string
		ExpectedDBPort     string
		ExpectedDBName     string
		ExpectedDBUser     string
		ExpectedDBPassword string
	}{
		{
			Name: "All variables set",
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
		},
		{
			Name:               "No variables set",
			EnvVars:            map[string]string{},
			ExpectedDBHost:     "",
			ExpectedDBPort:     "",
			ExpectedDBName:     "",
			ExpectedDBUser:     "",
			ExpectedDBPassword: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			setEnvVars(tc.EnvVars)
			config.LoadDBConfig()
			checkDBConfig(t, tc)
			unsetEnvVars(tc.EnvVars)
		})
	}
}

func TestLoadAppConfig(t *testing.T) {
	// test cases
	cases := []struct {
		Name         string
		EnvVars      map[string]string
		ExpectedPort string
	}{
		{
			Name: "Port set",
			EnvVars: map[string]string{
				"SERVER_PORT": "8080",
			},
			ExpectedPort: "8080",
		},
		{
			Name:         "Port not set",
			EnvVars:      map[string]string{},
			ExpectedPort: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			setEnvVars(tc.EnvVars)
			config.LoadAppConfig()
			checkAppConfig(t, tc)
			unsetEnvVars(tc.EnvVars)
		})
	}
}

func TestLoadAuthConfig(t *testing.T) {
	// test cases
	cases := []struct {
		Name                       string
		EnvVars                    map[string]string
		ExpectedJWTIssuer          string
		ExpectedJWTAudience        string
		ExpectedJWTSecret          string
		ExpectedAccessTokenExpiry  time.Duration
		ExpectedRefreshTokenExpiry time.Duration
		ExpectedJWTCookieDomain    string
		ExpectedJWTCookiePath      string
		ExpectedJWTCookieName      string
	}{
		{
			Name: "All variables set",
			EnvVars: map[string]string{
				"JWT_ISSUER":           "test_issuer",
				"JWT_AUDIENCE":         "test_audience",
				"JWT_SECRET":           "test_secret",
				"ACCESS_TOKEN_EXPIRY":  "900",
				"REFRESH_TOKEN_EXPIRY": "86400",
				"JWT_COOKIE_DOMAIN":    "localhost",
				"JWT_COOKIE_PATH":      "/",
				"JWT_COOKIE_NAME":      "token",
			},
			ExpectedJWTIssuer:          "test_issuer",
			ExpectedJWTAudience:        "test_audience",
			ExpectedJWTSecret:          "test_secret",
			ExpectedAccessTokenExpiry:  900 * time.Second,
			ExpectedRefreshTokenExpiry: 86400 * time.Second,
			ExpectedJWTCookieDomain:    "localhost",
			ExpectedJWTCookiePath:      "/",
			ExpectedJWTCookieName:      "token",
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			setEnvVars(tc.EnvVars)
			config.LoadAuthConfig()
			checkAuthConfig(t, tc)
			unsetEnvVars(tc.EnvVars)
		})
	}
}

func setEnvVars(vars map[string]string) {
	for k, v := range vars {
		os.Setenv(k, v)
	}
}

func unsetEnvVars(vars map[string]string) {
	for k := range vars {
		os.Unsetenv(k)
	}
}

func checkDBConfig(t *testing.T, tc struct {
	Name               string
	EnvVars            map[string]string
	ExpectedDBHost     string
	ExpectedDBPort     string
	ExpectedDBName     string
	ExpectedDBUser     string
	ExpectedDBPassword string
}) {
	if config.DBHost != tc.ExpectedDBHost ||
		config.DBPort != tc.ExpectedDBPort ||
		config.DBName != tc.ExpectedDBName ||
		config.DBUser != tc.ExpectedDBUser ||
		config.DBPassword != tc.ExpectedDBPassword {
		t.Fatalf("Loaded DB config does not match expected config")
	}
}

func checkAppConfig(t *testing.T, tc struct {
	Name         string
	EnvVars      map[string]string
	ExpectedPort string
}) {
	if config.Port != tc.ExpectedPort {
		t.Fatalf("Loaded App config does not match expected config")
	}
}

func checkAuthConfig(t *testing.T, tc struct {
	Name                       string
	EnvVars                    map[string]string
	ExpectedJWTIssuer          string
	ExpectedJWTAudience        string
	ExpectedJWTSecret          string
	ExpectedAccessTokenExpiry  time.Duration
	ExpectedRefreshTokenExpiry time.Duration
	ExpectedJWTCookieDomain    string
	ExpectedJWTCookiePath      string
	ExpectedJWTCookieName      string
}) {
	if config.JWTIssuer != tc.ExpectedJWTIssuer ||
		config.JWTAudience != tc.ExpectedJWTAudience ||
		config.JWTSecret != tc.ExpectedJWTSecret ||
		config.AccessTokenExpiry != tc.ExpectedAccessTokenExpiry ||
		config.RefreshTokenExpiry != tc.ExpectedRefreshTokenExpiry ||
		config.JWTCookieDomain != tc.ExpectedJWTCookieDomain ||
		config.JWTCookiePath != tc.ExpectedJWTCookiePath ||
		config.JWTCookieName != tc.ExpectedJWTCookieName {
		t.Fatalf("Loaded Auth config does not match expected config")
	}
}
