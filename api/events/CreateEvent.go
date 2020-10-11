package events

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/db"
)

func CreateEvent(c *fiber.Ctx) error {
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
	// token payload
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	tokenID := claims["id"].(string)
	// adds the user_id to the object
	body.UserID = tokenID
	// creates a new event
	created, err := prisma.Event.CreateOne(
		db.Event.Title.Set(body.Title),
		db.Event.Img.Set(body.Img),
		db.Event.Date.Set(body.Date),
		db.Event.Type.Set(body.Type),
		db.Event.Max.Set(body.Max),
		db.Event.HostName.Set(body.HostName),
		db.Event.HostURL.Set(body.HostURL),
		db.Event.Link.Set(body.Link),
		db.Event.UserID.Set(tokenID),
	).Exec(ctx)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	// Response
	return c.JSON(created)
}
