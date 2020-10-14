package attends

import (
	"context"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/db"
)

func Attend(c *fiber.Ctx) error {
	type FindRes struct {
		AttendsNumber int
	}
	var findResponse FindRes
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
	// token payload
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	tokenID := claims["id"].(string)
	// find event
	find, err := prisma.Event.FindOne(
		db.Event.ID.Equals(EventID),
	).Exec(ctx)
	findResponse = FindRes{
		AttendsNumber: find.AttendsNumber,
	}
	// create attendance
	create, err := prisma.Attends.CreateOne(
		db.Attends.EventID.Set(EventID),
		db.Attends.UserID.Set(tokenID),
	).Exec(ctx)
	// update event attends
	update, err := prisma.Event.FindOne(
		db.Event.ID.Equals(EventID),
	).Update(
		db.Event.AttendsNumber.Set(findResponse.AttendsNumber + 1),
	).Exec(ctx)
	log.Printf(update.ID)
	// response
	return c.JSON(create)
}
