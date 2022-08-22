package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/ipron-cloud/call-app/pkg/config"
	"gitlab.com/ipron-cloud/call-app/pkg/middleware"
)

func main() {
	// Default Fiber Config
	config := config.FiberConfig()

	// Define a new Fiber app with config
	app := fiber.New(config)

	// Server files from multiple directory
	app.Static("/views", "./views")

	// Set Middleware.
	middleware.FiberMiddleware(app)

	// Render index template
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{}, "layouts/main")
	})

	// Listen app
	log.Fatal(app.Listen(":3000"))
}
