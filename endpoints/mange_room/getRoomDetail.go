package mange_room

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func GetRoomDetail(c *fiber.Ctx) error {
	roomId := c.Query("roomId")
	if roomId == "" {
		return &response.GenericError{
			Message: "roomId is missing from query parameters",
			Err:     nil,
		}
	}
	var room []model.Room

	if result := mysql.Gorm.Where("Id  = ?", roomId).First(&room); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get room",
			Err:     result.Error,
		}
	}

	return c.JSON(room)
}