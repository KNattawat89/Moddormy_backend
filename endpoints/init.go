package endpoints

import (
	"Moddormy_backend/endpoints/upload"
	"github.com/gofiber/fiber/v2"
)

func Register(router fiber.Router) {

	uploadGroup := router.Group("upload/")
	uploadGroup.Post("dorm", upload.Dorming)
	uploadGroup.Post("room", upload.Rooming)

}
