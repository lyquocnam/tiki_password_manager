package app

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	FileName = "password.txt"
	DemoUser = "tiki"
	DemoPassword = "Tik@19"
)

func RemoveFileIfExist() error {
	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		return nil
	}
	return os.Remove(FileName)
}

func LoadDataFromFile() error {
	// check file exist


	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		// create new file
		User = &PasswordManager{
			Username: DemoUser,
		}

		err := User.Encrypt(DemoPassword)
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
	if len(username) < 3 {
		return errors.New(`Username can not less than 3`)
	} else if ok, err := User.ValidatePassword(password); err != nil && !ok {
		return err
	} else if username == User.Username {
		return errors.New(`User already exist`)
	}

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
		return errors.New(`Current password is not true`)
	}

	if _, err := User.ValidatePassword(newPassword); err != nil {
		return err
	}

	err := User.Encrypt(newPassword)
	if err != nil {
		return err
	}

	return SaveToFile()
}