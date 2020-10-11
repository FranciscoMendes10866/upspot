package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/attends"
	"github.com/upspot-api/guards"
)

func AttendsRouter(app *fiber.App) {
	// router prefix
	api := app.Group("/api/v1/attends")
	// routes
	api.Post("/:id", guards.Protected(), attends.Attend)
}
