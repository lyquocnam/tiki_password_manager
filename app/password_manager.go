package app

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type PasswordManager struct {
	Username          string `gorm:"unique; not null"`
	EncryptedPassword string `gorm:"not null" json:"-"`
}

var User *PasswordManager

func (u *PasswordManager) Encrypt(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.EncryptedPassword = string(hashed)
	return nil
}

func (u *PasswordManager) VerifyPassword(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(plainPassword))
	if err != nil {
		return false
	}
	return true
}

func (u *PasswordManager) ValidatePassword(password string) (bool, error) {
	if len(password) == 0 {
		return false, errors.New("The password can not be empty")
	} else if len(password) > 6 {
		return false, errors.New("The password can not be greater than 6")
	} else if matched, _ := regexp.MatchString(`[\n# $&:\n\t]`, password); matched {
		return false, errors.New("The password can not contain any white space")
	} else if matched, _ := regexp.MatchString(`^.*([a-z]).*$`, password); !matched {
		return false, errors.New("The password must contain at least one uppercase and at least one lowercase letter")
	} else if matched, _ := regexp.MatchString(`^.*([A-Z]).*$`, password); !matched {
		return false, errors.New("The password must contain at least one uppercase and at least one lowercase letter")
	} else if matched, _ := regexp.MatchString(`^.*([\d]).*$`, password); !matched {
		return false, errors.New("The password must have at least one digit and symbol")
	} else if matched, _ := regexp.MatchString(`^.*([!@#\$%\^&]).*$`, password); !matched {
		return false, errors.New("The password must have at least one digit and symbol")
	}
	//} else if matched, _ := regexp.MatchString("(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z])", password); !matched {
	//	print("matched: ", matched)
	//	return false, errors.New("The password must contain at least one uppercase and at least one lowercase letter")
	//}

	return true, nil
	// matched, _ := regexp.MatchString(`^(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z])(?=.*\d)(?=.*[!@#\$%\^&\*]).{3,6}$`, password)
	// return matched
}
