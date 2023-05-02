package mange_room

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func GetDormRoom(c *fiber.Ctx) error {

	dormId := c.Query("dormId")
	if dormId == "" {
		return &response.GenericError{
			Message: "dormId is missing from query parameters",
			Err:     nil,
		}
	}

	var room []model.Room

	if result := mysql.Gorm.Where("dorm_id  = ?", dormId).Find(&room); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get room",
			Err:     result.Error,
		}
	}

	return c.JSON(room)
}
