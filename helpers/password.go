package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pwd, 10)
}

func ComparePasswords(hashed, pwd []byte) error {
	return bcrypt.CompareHashAndPassword(hashed, pwd)
}
