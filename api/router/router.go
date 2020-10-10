package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/users"
)

func SetupRouter(app *fiber.App) {
	// router prefix
	api := app.Group("/api/v1")
	// auth routes
	auth := api.Group("/auth")
	auth.Post("/", users.SignUp)
}
