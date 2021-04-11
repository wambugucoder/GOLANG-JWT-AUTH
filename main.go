package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang_auth/database"
	"golang_auth/routes"
	"log"
	"time"
)

func main() {

	//INITIALIZE APP
	app := fiber.New()

	//USE CORS
	app.Use(cors.New())

	//DATABASE CONFIG
	database.ConnectDB()

	//SETUP ROUTES
	routes.SetupApiRoutes(app)

	//CACHE
	app.Use(cache.New(cache.Config{Next: func(c *fiber.Ctx) bool {
		return c.Query("refresh") == "true"
	},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	//RATE LIMITER
	app.Use(limiter.New(limiter.Config{
		Max:        5,
		Expiration: 10 * time.Second,
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.IP()
		},
		LimitReached: func(ctx *fiber.Ctx) error {
			return ctx.SendStatus(fiber.StatusTooManyRequests)
		},
	}))

	//COMPRESSIONS
	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))

	//404 NOT FOUND
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(400)
	})

	log.Fatal(app.Listen(":3000"))
}
