package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/users"
)

func AuthRouter(app *fiber.App) {
	// router prefix
	api := app.Group("/api/v1/auth")
	// routes
	api.Post("/sign_up", users.SignUp)
	api.Post("/sign_in", users.SignIn)
}
