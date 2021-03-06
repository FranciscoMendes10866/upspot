package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/events"
	"github.com/upspot-api/guards"
)

func EventsRouter(app *fiber.App) {
	// router prefix
	api := app.Group("/api/v1/events")
	// routes
	api.Post("/", guards.Protected(), events.CreateEvent)
	api.Post("/get_all", events.FindEvents)
	api.Get("/:id", events.FindEvent)
	api.Delete("/:id", guards.Protected(), events.DeleteEvent)
	api.Put("/:id", guards.Protected(), events.UpdateEvent)
}
