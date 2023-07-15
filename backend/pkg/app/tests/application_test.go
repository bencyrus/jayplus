package tests

import (
	"backend/pkg/app"
	"testing"
)

func TestNewApplication(t *testing.T) {
	_, err := app.NewApplication()
	if err != nil {
		t.Errorf("failed to setup app: %s", err.Error())
	}
}
