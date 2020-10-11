package events

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/db"
)

func FindEvents(c *fiber.Ctx) error {
	// Prisma Client
	prisma := db.NewClient()
	err := prisma.Connect()
	if err != nil {
		return err
	}
	defer func() {
		err := prisma.Disconnect()
		if err != nil {
			panic(err)
		}
	}()
	ctx := context.Background()
	// user input
	// User Input
	body := new(Event)
	c.BodyParser(body)
	cityInput := body.City
	// database query
	query, err := prisma.Event.FindMany(
		db.Event.City.Equals(cityInput),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	// response
	return c.JSON(query)
}
