package events

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/db"
)

func DeleteEvent(c *fiber.Ctx) error {
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
	// delete method
	event, err := prisma.Event.FindOne(
		db.Event.ID.Equals(EventID),
	).Delete().Exec(ctx)
	if err != nil {
		panic(err)
	}
	// response
	return c.JSON(event.ID)
}
