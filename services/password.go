package services

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)

}
