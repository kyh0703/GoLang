package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SwaggerRoute(app *fiber.App) {
	// Create routes groups.
	route := app.Group("/swagger")

	// Routes for GET Method
	route.Get("*", swagger.HandlerDefault)
}
