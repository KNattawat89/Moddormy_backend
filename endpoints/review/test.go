package review

import "github.com/gofiber/fiber/v2"

func Test(c *fiber.Ctx) error {
	return c.SendString("hello from review route")
}
