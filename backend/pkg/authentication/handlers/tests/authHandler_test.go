package tests

import (
	"backend/pkg/authentication/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
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
			handler := http.HandlerFunc(handlers.LoginHandler)

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
