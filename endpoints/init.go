package endpoints

import (
	"github.com/gofiber/fiber/v2"

	"Moddormy_backend/endpoints/upload"
)

func Load(app *fiber.App) {
	api := app.Group("api/")

	uploadGroup := api.Group("upload/")
	uploadGroup.Post("base", upload.Base64)

}
