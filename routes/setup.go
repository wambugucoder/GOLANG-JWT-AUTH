package routes

import "github.com/gofiber/fiber/v2"

//setUpApiRoutes->Connect all rest Apis
func SetupApiRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/", hello)

}

func hello(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello")
}
