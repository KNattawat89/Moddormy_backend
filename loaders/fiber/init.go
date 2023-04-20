package fiber

import (
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/config"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func Init() {
	// Initialize fiber instance
	app = fiber.New(fiber.Config{
		ErrorHandler:  errorHandler,
		Prefork:       false,
		StrictRouting: true,
		ServerHeader:  config.C.ServerHeader,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  5 * time.Second,
	})

	// Register root endpoint
	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(response.InfoResponse{
			Success: true,
			Message: "CHOUXCREAM_API_ROOT",
		})
	})

	// Register API endpoints
	apiGroup := app.Group("api/")

	//apiGroup.Use(middlewares.Limiter)
	//apiGroup.Use(middlewares.Cors)
	//apiGroup.Use(middlewares.Recover)

	//endpoints.Register(apiGroup)

	apiGroup.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// Register not found handler
	app.Use(notfoundHandler)

	// Startup
	err := app.Listen(config.C.BackAddress)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
