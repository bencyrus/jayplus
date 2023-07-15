package tests

import (
	"backend/pkg/app"
	"testing"
)

func TestSetupApp(t *testing.T) {
	err := app.SetupApp()
	if err != nil {
		t.Errorf("failed to setup app: %s", err.Error())
	}
}
