package attends

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/db"
)

func Attend(c *fiber.Ctx) error {
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
	// create attendance
	create, err := prisma.Attends.CreateOne(
		db.Attends.EventID.Set(EventID),
		db.Attends.UserID.Set(tokenID),
	).Exec(ctx)
	// response
	return c.JSON(create)
}
