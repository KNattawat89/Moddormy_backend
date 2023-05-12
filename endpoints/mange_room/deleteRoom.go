package mange_room

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func DeleteRoom(c *fiber.Ctx) error {

	dormId := c.Query("dormId")

	if dormId == "" {
		return &response.GenericError{
			Message: "dormId is missing from query parameters",
			Err:     nil,
		}
	}

	var room []model.Room
	if result := mysql.Gorm.Where("dorm_id = ?", dormId).Find(&room); result.Error != nil {
		return &response.GenericError{
			Message: "Room not found",
			Err:     result.Error,
		}
	}

	for i := 0; i < len(room); i++ {
		roomId := room[i].Id
		var roomImages []model.RoomImage
		if result := mysql.Gorm.Where(roomId, roomId).Find(&roomImages); result.Error != nil {
			return &response.GenericError{
				Message: "RoomImage not found",
				Err:     result.Error,
			}
		}
		if result := mysql.Gorm.Delete(&roomImages); result.Error != nil {
			return &response.GenericError{
				Message: "Unable to delete roomImage",
				Err:     result.Error,
			}
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
