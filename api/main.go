package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/router"
)

func main() {
	app := fiber.New()

	router.SetupRouter(app)

	app.Listen(":3333")
}
