package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(passwod string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwod), bcrypt.DefaultCost)
	
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	
	if err != nil {
		return false, err
	}

	return true, nil
}