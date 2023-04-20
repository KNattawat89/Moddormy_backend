package endpoints

import (
	"Moddormy_backend/endpoints/test"
	"github.com/gofiber/fiber/v2"
)

func Register(router fiber.Router) {
	testza := router.Group("test/")
	testza.Get("hello", test.Hello)
}
