package tests

import (
	"backend/pkg/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	a, err := app.NewApplication()
	if err != nil {
		t.Fatalf("failed to create application: %s", err.Error())
	}

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("failed to create request: %s", err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.Router.ServeHTTP)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `test handler`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
