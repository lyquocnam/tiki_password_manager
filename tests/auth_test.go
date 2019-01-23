package tests

import (
	"testing"
	"github.com/lyquocnam/tiki_password_manager/app"
)

func TestLoadData(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	if app.User == nil {
		t.Errorf(`User can not be nil`)
	}
}


func TestEncryptPassword(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	err = app.User.Encrypt(app.DemoPassword)
	if err != nil {
		t.Error(err)
	}
}

func TestLogin(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	err = app.Login(app.DemoUser, app.DemoPassword)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateUser(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	var user = "tiki2"
	var password = "Tik!12"

	err = app.CreateUser(user, password)
	if err != nil {
		t.Error(err)
	}

	err = app.Login(user, password)
	if err != nil {
		t.Error(err)
	}

	err = app.RemoveFileIfExist()
	if err != nil {
		t.Error(err)
	}
}

func TestChangePassword(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	var newPassword = "Tik!12"

	err = app.ChangePassword(app.DemoPassword, newPassword)
	if err != nil {
		t.Error(err)
	}

	err = app.Login(app.DemoUser, newPassword)
	if err != nil {
		t.Error(err)
	}

	err = app.RemoveFileIfExist()
	if err != nil {
		t.Error(err)
	}
}


func TestValidatePasswordShouldBeOK(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	ok, err := app.User.ValidatePassword(app.DemoPassword)
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf(`VerifyPassword should be return true`)
	}
}

func TestValidatePasswordShouldBeEmpty(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	_, err = app.User.ValidatePassword("")
	if err != nil {
		if err.Error() != "The password can not be empty" {
			t.Error(err)
		}
	}
}

func TestValidatePasswordShouldBeLessThan6(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	_, err = app.User.ValidatePassword("1234567")
	if err != nil {
		if err.Error() != "The password can not be greater than 6" {
			t.Error(err)
		}
	}
}

func TestValidatePasswordShouldBeContainWhiteSpace(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	_, err = app.User.ValidatePassword("123 67")
	if err != nil {
		if err.Error() != "The password can not contain any white space" {
			t.Error(err)
		}
	}
}

func TestValidatePasswordShouldBeContainOneUpperCase(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	_, err = app.User.ValidatePassword("n23456")
	if err != nil {
		if err.Error() != "The password must contain at least one uppercase and at least one lowercase letter" {
			t.Error(err)
		}
	}
}

func TestValidatePasswordShouldBeContainOneLowerCase(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	_, err = app.User.ValidatePassword("N23456")
	if err != nil {
		if err.Error() != "The password must contain at least one uppercase and at least one lowercase letter" {
			t.Error(err)
		}
	}
}

func TestValidatePasswordShouldBeContainOneDigit(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	_, err = app.User.ValidatePassword("lyNamt")
	if err != nil {
		if err.Error() != "The password must have at least one digit and symbol" {
			t.Error(err)
		}
	}
}

func TestValidatePasswordShouldBeContainOneSymbol(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	_, err = app.User.ValidatePassword("lyNam3")
	if err != nil {
		if err.Error() != "The password must have at least one digit and symbol" {
			t.Error(err)
		}
	}
}

func TestVerifyPassword(t *testing.T) {
	err := app.LoadDataFromFile()
	if err != nil {
		t.Error(err)
	}

	ok := app.User.VerifyPassword(app.DemoPassword)
	if !ok {
		t.Errorf(`VerifyPassword should be return true`)
	}
}
