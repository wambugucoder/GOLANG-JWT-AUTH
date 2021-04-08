package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang_auth/database"
	"golang_auth/routes"
	"log"
)

func main() {

	app := Setup()

	log.Fatal(app.Listen(":3000"))

}
func Setup() *fiber.App {
	//INITIALIZE APP
	app := fiber.New()

	//USE CORS
	app.Use(cors.New())

	//DATABASE CONFIG
	database.ConnectDB()

	//SETUP ROUTES
	routes.SetupApiRoutes(app)

	//404 NOT FOUND
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(400)
	})

	return app
}
