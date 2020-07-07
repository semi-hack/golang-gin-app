package services

import (
	"golang.org/x/crypto/bcrypt"
)


// PasswordCompare ...
func PasswordCompare(password []byte, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	return err
}