package utils


import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)


func PasswordMatches(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		switch {
		// 不匹配
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		// 其它
		default:
			return false, err
		}
	}

	return true, nil
}

func PasswordMatchesV1(hash, password string) (int, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, nil 
		} else {
			return 0, err
		}
	}

	return 1, nil
}

func GenerateHashFromPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

func VerifyHashAndPassword(hashedPassword, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

