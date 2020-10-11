package attends

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/db"
)

func NotAttend(c *fiber.Ctx) error {
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
	eventID := c.Params("id")
	// token payload
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	tokenID := claims["id"].(string)
	// delete method
	delete, err := prisma.Attends.FindMany(
		db.Attends.EventID.Equals(eventID),
		db.Attends.UserID.Equals(tokenID),
	).Delete().Exec(ctx)
	if err != nil {
		panic(err)
	}
	// response
	return c.JSON(delete)
}
