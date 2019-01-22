package tests

import (
	"testing"
	"tiki/app"
)

func TestLogin(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	err = app.Login("lynam", "Lyn@m3")
	if err != nil {
		t.Error(err)
	}
}
