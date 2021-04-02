package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"golang_auth/configs"
	"golang_auth/routes"
	"log"
)

func main() {
	//INITIALIZE APP
	app := fiber.New()

	//USE CORS
	app.Use(cors.New())

	//DATABASE CONFIG
	configs.ConnectDatabase()

	//SETUP ROUTES
	routes.SetupApiRoutes(app)

	//404 NOT FOUND
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(400)
	})

	log.Fatal(app.Listen(":3000"))

}
