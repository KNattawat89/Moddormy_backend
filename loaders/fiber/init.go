package fiber

import (
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/config"
	"Moddormy_backend/utils/logger"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/gofiber/fiber/v2"

	"Moddormy_backend/endpoints"
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
			Message: "Moddormy_API_ROOT",
		})
	})

	// Register API endpoints
	apiGroup := app.Group("api/")

	//apiGroup.Use(middlewares.Limiter)
	//apiGroup.Use(middlewares.Cors)
	//apiGroup.Use(middlewares.Recover)

	endpoints.Register(apiGroup)

	//apiGroup.Get("/hello", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello, World 👋!")
	//})
	app.Static("images/", "./images")

	// Register not found handler
	app.Use(notfoundHandler)

	// Startup
	err := app.Listen(config.C.BackAddress)
	if err != nil {
		logger.Log(logrus.Fatal, err.Error())
	}
}

//import (
//	"time"
//
//	"github.com/gofiber/fiber/v2"
//
//	"Moddormy_backend/endpoints"
//	"Moddormy_backend/loaders/storage"
//	"Moddormy_backend/types"
//	"Moddormy_backend/utils/config"
//	"Moddormy_backend/utils/wrapper"
//)
//
//var App *fiber.App
//
//func Init() {
//	// Initialize fiber instance
//	App = fiber.New(fiber.Config{
//		Prefork:       false,
//		StrictRouting: true,
//		ReadTimeout:   30 * time.Second,
//		WriteTimeout:  30 * time.Second,
//		BodyLimit:     512 * 1024 * 1024,
//		ErrorHandler:  defaultErrorHandler,
//	})
//
//	// Import middlewares
//	//App.Use(corsMiddleware)
//	//App.Use(recoverMiddleware)
//
//	// Import static files
//	App.Static("/files", storage.Dir)
//
//	// Load endpoints
//	App.Get("/", func(c *fiber.Ctx) error {
//		return &types.PassError{Message: "API_ROOT"}
//	})
//
//	endpoints.Load(App)
//
//	//App.Use(notFoundMiddleware)
//
//	// Startup
//	err := App.Listen(config.C.BackAddress)
//	if err != nil {
//		wrapper.Fatal(err.Error())
//	}
//}
