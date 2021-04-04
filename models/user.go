package models

// User struct
type User struct {
	Base
	Username string `json:"username" gorm:"unique" validate:"required,min=6,max=32"`
	Email    string `json:"email" gorm:"unique" validate:"required,email,min=6,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}
