package authentication

import "github.com/gofiber/fiber/v2"

func Test(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋! from authentication route")
}
