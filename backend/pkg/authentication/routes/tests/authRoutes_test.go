package tests

import (
	"backend/pkg/authentication/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestAuthRoutes(t *testing.T) {
	cases := []struct {
		name           string
		route          string
		method         string
		expectedStatus int
	}{
		{name: "Login", route: "/login", method: "GET", expectedStatus: http.StatusOK},
		{name: "Login with POST", route: "/login", method: "POST", expectedStatus: http.StatusMethodNotAllowed},
		{name: "Non-existing route", route: "/non_existing", method: "GET", expectedStatus: http.StatusNotFound},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := mux.NewRouter()
			routes.AuthRoutes(r)

			req, _ := http.NewRequest(tc.method, tc.route, nil)
			response := httptest.NewRecorder()
			r.ServeHTTP(response, req)

			if status := response.Code; status != tc.expectedStatus {
				t.Errorf("Test failed, expected: '%v', got:  '%v'", tc.expectedStatus, status)
			}
		})
	}
}
