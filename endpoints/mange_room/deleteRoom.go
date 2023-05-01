package mange_room

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func DeleteRoom(c *fiber.Ctx) error {

	roomId := c.Query("roomId")

	if roomId == "" {
		return &response.GenericError{
			Message: "roomId is missing from query parameters",
			Err:     nil,
		}
	}

	var room model.Room
	if result := mysql.Gorm.Where("Id = ?", roomId).First(&room); result.Error != nil {
		return &response.GenericError{
			Message: "Room not found",
			Err:     result.Error,
		}
	}

	if result := mysql.Gorm.Delete(&room); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to delete room",
			Err:     result.Error,
		}
	}

	return c.JSON(&response.GenericSuccess{
		Message: "Room deleted successfully",
	})

}
