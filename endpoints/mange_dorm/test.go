package mange_dorm

import "github.com/gofiber/fiber/v2"

func Test(c *fiber.Ctx) error {
	return c.SendString("Hello from manage dorm route")
}
