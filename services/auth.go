package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang_auth/database"
	"golang_auth/models"
)

type RegisterErrors struct {
	FailedField string
	Tag         string
	Value       string
}

func RegisterUser(ctx *fiber.Ctx) error {
	userDetails := new(models.User)

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

	results := database.DB.Create(&userDetails)

	return ctx.JSON(fiber.Map{
		"error":   false,
		"general": results,
	})
}

func ValidateRegistration(user models.User) []RegisterErrors {
	var errors []RegisterErrors

	validate := validator.New()

	//validate struct user
	err := validate.Struct(user)
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
