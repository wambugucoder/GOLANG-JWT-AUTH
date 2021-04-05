package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang_auth/database"
	"golang_auth/models"
	"golang_auth/repository"
)

type RegistrationInput struct {
	Username string `json:"username" validate:"required,min=6,max=32"`
	Email    string `json:"email"    validate:"required,email,min=6,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}
type LoginInputs struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}
type AuthErrors struct {
	FailedField string
	Tag         string
	Value       string
}

//RegisterUser->RegisterUser details
func RegisterUser(ctx *fiber.Ctx) error {
	userDetails := new(RegistrationInput)

	err := ctx.BodyParser(userDetails)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}
	if !repository.DoesEmailExist(userDetails.Email) {
		errors := ValidateAuth(*userDetails)
		if errors != nil {
			return ctx.JSON(errors)

		}

		//ENCRYPT PASSWORD FIRST
		userDetails.Password = GenerateHash(userDetails.Password)

		userInfo := &models.User{
			Username: userDetails.Username,
			Email:    userDetails.Email,
			Password: userDetails.Password,
		}

		results := database.DB.Create(userInfo)

		return ctx.JSON(fiber.Map{
			"error":   false,
			"general": results,
		})

	}
	return ctx.JSON(fiber.Map{
		"error":   true,
		"general": "Email already exists",
	})

}
func LoginUser(ctx *fiber.Ctx) error {
	logindetails := new(LoginInputs)

	err := ctx.BodyParser(logindetails)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	errors := ValidateLogin(*logindetails)
	if errors != nil {
		return ctx.JSON(errors)
	}
	if repository.DoesEmailExist(logindetails.Email) {
		//FETCH USER DETAILS TO COMPARE PASSWORDS
		user, _ := repository.GetUserDetailsByEmail(logindetails.Email)

		if !ComparePasswords(logindetails.Password, user.Password) {
			return ctx.JSON(fiber.Map{
				"error":   true,
				"general": "Passwords don't Match",
			})
		}
		//GENERATE JWT TOKEN
		credentials := GenerateJwtToken(user)
		return ctx.JSON(fiber.Map{
			"error": false,
			"token": "Bearer " + credentials,
		})

	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Email doesn't exist",
	})
}

func ValidateAuth(input RegistrationInput) []AuthErrors {
	var errors []AuthErrors

	validate := validator.New()

	//validate struct user
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var elements AuthErrors

			elements.FailedField = err.StructNamespace()
			elements.Tag = err.Tag()
			elements.Value = err.Param()

			errors = append(errors, elements)
		}
	}
	return errors
}
func ValidateLogin(input LoginInputs) []AuthErrors {
	var errors []AuthErrors

	validate := validator.New()

	//validate struct user
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var elements AuthErrors

			elements.FailedField = err.StructNamespace()
			elements.Tag = err.Tag()
			elements.Value = err.Param()

			errors = append(errors, elements)
		}
	}
	return errors
}
