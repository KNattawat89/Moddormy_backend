package test

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	//claim := c.Locals("user").(*jwt.Token).Claims.(*common.UserClaims)
	return c.SendString("Welcome Ja")
}
