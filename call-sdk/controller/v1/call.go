package controller

import "github.com/gofiber/fiber/v2"

type CallController struct {
	app *fiber.App
}

func NewCallController() *CallController {
	ctrl := &CallController{
		app: fiber.New(),
	}
	ctrl.Urls()
	return ctrl
}

func (ctrl *CallController) Fiber() *fiber.App {
	return ctrl.app
}

func (ctrl *CallController) MakeCall(c *fiber.Ctx) error {
	return c.SendString("make-call")
}

func (ctrl *CallController) HoldCall(c *fiber.Ctx) error {
	return c.SendString("hold-call")
}

func (ctrl *CallController) UnHoldCall(c *fiber.Ctx) error {
	return c.SendString("unhold-call")
}

func (ctrl *CallController) Urls() {
	ctrl.app.Get("/make-call", ctrl.MakeCall)
	ctrl.app.Get("/hold-call", ctrl.HoldCall)
	ctrl.app.Get("/unhold-call", ctrl.UnHoldCall)
}
