package controllers

import "github.com/gofiber/fiber/v2"

/**
* @RequestMapping("/api/v1/call")
 */
type callController struct {
	app *fiber.App
}

func NewCallController(app *fiber.App) *callController {
	ctrl := &callController{
		app: app,
	}
	ctrl.Urls()
	return ctrl
}

/**
* @ RequestMapping("/api/v1/call/make-call")
 */
func (ctrl *callController) MakeCall(c *fiber.Ctx) error {
	return c.SendString("make-call")
}

/**
* @ RequestMapping("/api/v1/call/hold-call")
 */
func (ctrl *callController) HoldCall(c *fiber.Ctx) error {
	return c.SendString("hold-call")
}

/**
* @ RequestMapping("/api/v1/call/unhold-call")
 */
func (ctrl *callController) UnHoldCall(c *fiber.Ctx) error {
	return c.SendString("unhold-call")
}

func (ctrl *callController) Urls() {
	ctrl.app.Get("/make-call", ctrl.MakeCall)
	ctrl.app.Get("/hold-call", ctrl.HoldCall)
	ctrl.app.Get("/unhold-call", ctrl.UnHoldCall)
}
