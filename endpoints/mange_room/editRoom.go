package mange_room

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func UpdateRoom(c *fiber.Ctx) error {
	roomId := c.Query("roomId")
	var room model.Room

	// check if dorm with specified ID exists
	if err := mysql.Gorm.Where("Id = ?", roomId).First(&room).Error; err != nil {
		return &response.GenericError{
			Message: "Room not found",
			Err:     err,
		}
	}

	// update dorm record
	if err := c.BodyParser(&room); err != nil {
		return &response.GenericError{
			Message: "Failed to parse request body",
			Err:     err,
		}
	}

	if err := mysql.Gorm.Save(&room).Error; err != nil {
		return &response.GenericError{
			Message: "Failed to update room",
			Err:     err,
		}
	}

	return c.JSON(room)
}
