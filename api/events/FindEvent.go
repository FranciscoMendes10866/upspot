package events

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/db"
)

func FindEvent(c *fiber.Ctx) error {
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
	// Event ID
	EventID := c.Params("id")
	// database query
	query, err := prisma.Event.FindOne(
		db.Event.ID.Equals(EventID),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	// response
	return c.JSON(query)
}
