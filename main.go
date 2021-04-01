package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	//INITIALIZE APP
	app := fiber.New()

	//USE CORS
	app.Use(cors.New())

	// EXAMPLE API
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})

	//404 NOT FOUND
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(400)
	})

	log.Fatal(app.Listen(":3000"))

}
