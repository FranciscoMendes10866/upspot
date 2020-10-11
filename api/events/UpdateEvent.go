package events

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/db"
)

func UpdateEvent(c *fiber.Ctx) error {
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
	// User Input
	body := new(Event)
	c.BodyParser(body)
	// Event ID
	EventID := c.Params("id")
	// update method
	event, err := prisma.Event.FindOne(
		db.Event.ID.Equals(EventID),
	).Update(
		db.Event.Title.Set(body.Title),
		db.Event.Img.Set(body.Img),
		db.Event.Date.Set(body.Date),
		db.Event.Type.Set(body.Type),
		db.Event.Max.Set(body.Max),
		db.Event.HostName.Set(body.HostName),
		db.Event.HostURL.Set(body.HostURL),
		db.Event.Link.Set(body.Link),
	).Exec(ctx)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	// Response
	return c.JSON(event)
}
