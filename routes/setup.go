package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang_auth/services"
)

//setUpApiRoutes->Connect all rest Apis
func SetupApiRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/hello", hello)

	//AUTH ROUTES
	auth := app.Group("/api/v1/auth")
	auth.Post("/register", services.RegisterUser)
	auth.Post("/login", services.LoginUser)
	auth.Post("/trial", services.CreateUserAndTweets)
	auth.Get("/fetch", services.FetchTweets)

}

func hello(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello")
}
