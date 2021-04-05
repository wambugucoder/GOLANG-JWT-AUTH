package repository

import (
	"github.com/jinzhu/gorm"
	"golang_auth/database"
	"golang_auth/models"
)

func DoesEmailExist(email string) bool {
	var user models.User
	err := database.DB.Where(&models.User{Email: email}).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false
		}
	}
	return true
}

func GetUserDetailsByEmail(email string) (*models.User, error) {
	var user models.User

	err := database.DB.Where(&models.User{Email: email}).Find(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil

}
