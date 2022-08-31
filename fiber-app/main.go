package main

import (
	"log"

	"fiber/pkg/config"
	"fiber/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Default Fiber Config
	cfg := config.FiberConfig()

	// Define a new Fiber app with config
	app := fiber.New(cfg)

	// Server files from multiple directory
	app.Static("/static", "./static")

	// Set Middleware.
	middleware.FiberMiddleware(app)

	// Render index template
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":   "Hey!",
			"Message": "This is the index templates.",
		})
	})

	// Listen app
	log.Fatal(app.Listen(":" + config.Env.Port))
}
