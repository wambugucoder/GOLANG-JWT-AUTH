package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RegisterData struct {
	Username string `validate:"required,min=6,max=32"`
	Email    string `validate:"required,email,min=6,max=32"`
	Password string `validate:"required,min=6,max=32"`
}

type RegisterErrors struct {
	FailedField string
	Tag         string
	Value       string
}

func RegisterUser(ctx *fiber.Ctx) error {
	registerDetails := new(RegisterData)

	err := ctx.BodyParser(registerDetails)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})

	}
	errors := ValidateRegistration(*registerDetails)
	if errors != nil {
		return ctx.JSON(errors)

	}
	//REGISTER USERS
	return ctx.JSON(registerDetails)

}

func ValidateRegistration(data RegisterData) []RegisterErrors {
	var errors []RegisterErrors

	validate := validator.New()

	//validate struct user
	err := validate.Struct(data)
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
