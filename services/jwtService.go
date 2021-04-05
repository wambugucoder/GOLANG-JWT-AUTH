package services

import (
	"github.com/dgrijalva/jwt-go"
	"golang_auth/config"
	"golang_auth/models"
	"time"
)

func GenerateJwtToken(user *models.User) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.UUID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenstr, err := token.SignedString([]byte(config.Config("JWT_KEY")))
	if err != nil {
		panic(err)
	}
	return tokenstr

}
