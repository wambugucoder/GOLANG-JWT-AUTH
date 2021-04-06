package services

import (
	"github.com/dgrijalva/jwt-go"
	"golang_auth/config"
	"golang_auth/models"
	"log"
	"time"
)

func GenerateJwtToken(user *models.User) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.UUID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = user.Email

	tokenstr, err := token.SignedString([]byte(config.Config("JWT_KEY")))
	if err != nil {
		panic(err)
	}
	return tokenstr

}
func ExtractClaims(tokenstr string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("JWT_KEY")), nil
	})

	if token.Valid && err == nil {
		claims := token.Claims.(jwt.MapClaims)

		log.Println("Token has been verified")
		return claims, true

	}
	log.Println("Invalid token detected")
	return nil, false

}
