package users

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/upspot-api/db"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *fiber.Ctx) error {
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
	body := new(User)
	c.BodyParser(body)
	// stores the user input password value
	pass := body.Password
	// verifies if the user exists
	response, err := prisma.User.FindMany(
		db.User.Email.Equals(body.Email),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	// here we check if the passwords match
	match := bcrypt.CompareHashAndPassword([]byte(response[0].Password), []byte(pass))
	if match != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = response[0].ID
	// Generate encoded token and send it as response.
	loginToken, err := token.SignedString([]byte("ZKcpd63DOb"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	// Response
	return c.JSON(fiber.Map{
		"token": loginToken,
	})
}
