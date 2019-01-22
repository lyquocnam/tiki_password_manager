package app

import (
	"errors"
	"github.com/alecthomas/gometalinter/_linters/src/gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var FileName = "password.txt"

func LoadDataFromFile() error {
	// check file exist
	demoUser := "lynam"
	demoPassword := "Lyn@m2"

	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		// create new file
		User = &PasswordManager{
			Username: demoUser,
		}

		err := User.Encrypt(demoPassword)
		if err != nil {
			return err
		}

		return SaveToFile()
	} else {
		f, err := ioutil.ReadFile(FileName)
		if err != nil {
			return err
		}

		pm := &PasswordManager{}

		if err := yaml.Unmarshal(f, &pm); err != nil {
			return err
		}

		User = pm
		return nil
	}
}

func SaveToFile() error {
	fileContent, err := yaml.Marshal(&User)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(FileName, fileContent, 0644); err != nil {
		return err
	}
	return nil
}

func Login(username string, password string) error {
	if User.Username != username {
		return errors.New(`Username is not valid`)
	} else {
		if _, err := User.ValidatePassword(password); err != nil {
			return err
		}
		return nil
	}
}

func CreateUser(username string, password string) error {
	User = &PasswordManager{
		Username: username,
	}
	err := User.Encrypt(password)
	if err != nil {
		return err
	}

	return SaveToFile()
}

func ChangePassword(password string, newPassword string) error {
	if !User.VerifyPassword(password) {
		return errors.New(`Password is not valid`)
	}

	err := User.Encrypt(newPassword)
	if err != nil {
		return err
	}

	return SaveToFile()
}