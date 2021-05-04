package toolkit

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GeneratePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(passwordNotCheck string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordNotCheck))
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}
