package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")

	middleware := func(c *fiber.Ctx) error {
		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	}

	listHandler := func(c *fiber.Ctx) error {
		return c.SendString("test")
	}

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	api := app.Group("/api", middleware)
	v1 := api.Group("/v1", middleware)
	v1.Get("/list", listHandler)
	v1.Get("/user", listHandler)

	v2 := api.Group("/v2", middleware)
	v2.Get("/list", listHandler)
	v2.Get("/user", listHandler)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	// Render index template
	// 	return c.Render("index", fiber.Map{"Title": "Hello, world!"})
	// })

	// app.Get("/user/+", func(c *fiber.Ctx) error {
	// 	return c.SendString(c.Params("+"))
	// })

	// app.Get("/user/:name?", func(c *fiber.Ctx) error {
	// 	return c.SendString(c.Params("name"))
	// })

	// app.Get("/plantae/:genus.:species", func(c *fiber.Ctx) error {
	// 	fmt.Fprintf(c, "%s.%s\n", c.Params("genus"), c.Params("species"))
	// 	return nil // prunus.persica
	// })
	log.Fatal(app.Listen(":3000"))
}
