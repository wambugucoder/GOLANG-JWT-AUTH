package models

type User struct {
	Base
	username string `json:"username" gorm:"unique"`
	email    string `json:"email" gorm:"unique"`
	password string `json:"password"`
}
