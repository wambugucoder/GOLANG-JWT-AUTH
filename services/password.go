package services

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hash)

}

func ComparePasswords(raw string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw))
	if err != nil {
		return false
	}
	return true
}
