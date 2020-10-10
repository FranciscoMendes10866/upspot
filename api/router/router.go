package router

import "github.com/gofiber/fiber/v2"

func SetupRouter(app *fiber.App) {
	// router prefix
	api := app.Group("/api/v1")
	// auth routes
	auth := api.Group("/auth")
	auth.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Router working")
	})
}
