package models

type User struct {
	Base
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type UserError struct {
	Err      bool   `json:"err"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
