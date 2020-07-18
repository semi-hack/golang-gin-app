package services

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)


// PasswordCompare ...
func PasswordCompare(password []byte, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
    if err != nil {
		log.Println("error")
		return nil
	}
	return err
}