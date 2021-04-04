package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang_auth/database"
	"golang_auth/models"
)

type RegistrationInput struct {
	Username string `json:"username" validate:"required,min=6,max=32"`
	Email    string `json:"email"    validate:"required,email,min=6,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}
type RegisterErrors struct {
	FailedField string
	Tag         string
	Value       string
}

func RegisterUser(ctx *fiber.Ctx) error {
	userDetails := new(RegistrationInput)

	err := ctx.BodyParser(userDetails)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})

	}
	errors := ValidateRegistration(*userDetails)
	if errors != nil {
		return ctx.JSON(errors)

	}

	//ENCRYPT PASSWORD FIRST
	userDetails.Password = GenerateHash(userDetails.Password)

	deeds := &models.User{
		Username: userDetails.Username,
		Email:    userDetails.Email,
		Password: userDetails.Password,
	}

	results := database.DB.Create(deeds)

	return ctx.JSON(fiber.Map{
		"error":   false,
		"general": results,
	})
}

func ValidateRegistration(input RegistrationInput) []RegisterErrors {
	var errors []RegisterErrors

	validate := validator.New()

	//validate struct user
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var elements RegisterErrors

			elements.FailedField = err.StructNamespace()
			elements.Tag = err.Tag()
			elements.Value = err.Param()

			errors = append(errors, elements)
		}
	}
	return errors
}
