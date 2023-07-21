package tests

import (
	"backend/pkg/authentication"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLoginHandler(t *testing.T) {
	// Create an instance of Auth for testing
	auth := &authentication.Auth{
		Issuer:             "testIssuer",
		Audience:           "testAudience",
		Secret:             "testSecret",
		AccessTokenExpiry:  1 * time.Hour,
		RefreshTokenExpiry: 24 * time.Hour,
		CookieDomain:       "localhost",
		CookiePath:         "/",
		CookieName:         "refresh_token",
	}

	cases := []struct {
		name           string
		method         string
		expectedStatus int
		expectedOutput string
	}{
		{name: "GET", method: "GET", expectedStatus: http.StatusOK, expectedOutput: `login handler`},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "/login", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(auth.LoginHandler) // Use the auth instance

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.expectedStatus)
			}

			if rr.Body.String() != tc.expectedOutput {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tc.expectedOutput)
			}
		})
	}
}
