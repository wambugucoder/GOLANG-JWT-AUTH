package services

import (
	"github.com/gofiber/fiber/v2"
	"golang_auth/models"
	"golang_auth/repository"
)

type UserAndTweets struct {
	Username string `json:"username" validate:"required,min=6,max=32"`
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
	Content  string `json:"content" validate:"required,min=6,max=32"`
}

func CreateUserAndTweets(ctx *fiber.Ctx) error {
	userplustweets := new(UserAndTweets)

	err := ctx.BodyParser(userplustweets)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}
	tweetDetails := []models.Tweet{{
		Content: userplustweets.Content,
	},
	}
	usertweetDetails := &models.User{
		Username: userplustweets.Username,
		Email:    userplustweets.Email,
		Password: userplustweets.Password,
		Tweets:   tweetDetails,
	}
	ok := repository.SaveUserAndTweets(usertweetDetails)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "oops",
		})
	}
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"success": "user &tweet saved",
	})
}
func FetchTweets(ctx *fiber.Ctx) error {
	tweets, didItFind := repository.FindAllTweets()
	if !didItFind {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "oops",
		})
	}
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"tweets": tweets,
	})

}
