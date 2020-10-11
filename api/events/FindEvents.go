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
	body := new(Event)
	c.BodyParser(body)
	titleInput := body.Title
	cityInput := body.City
	// database query
	// queries the database looking for the given title and city
	complete, err := prisma.Event.FindMany(
		db.Event.Title.Equals(titleInput),
		db.Event.City.Equals(cityInput),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	// verifies if there are many items
	// if there is one, we will send it
	if len(complete) > 0 {
		return c.JSON(complete)
	}
	// if the array is empty
	// we will query only the city events
	query, err := prisma.Event.FindMany(
		db.Event.City.Equals(cityInput),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	// final query response
	return c.JSON(query)
}
