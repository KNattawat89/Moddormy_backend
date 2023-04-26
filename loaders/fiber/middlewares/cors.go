package middlewares

import (
	"Moddormy_backend/utils/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Cors = func() fiber.Handler {
	origins := ""
	for i, s := range config.C.Cors {
		origins += s
		if i < len(config.C.Cors) {
			origins += ", "
		}
	}

	config := cors.Config{
		AllowOrigins:     origins,
		AllowCredentials: true,
	}

	return cors.New(config)
}()

// why we need to have "()" at the end?
