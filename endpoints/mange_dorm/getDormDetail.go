package mange_dorm

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func GetDormDetail(c *fiber.Ctx) error {
	dormId := c.Query("dormId")
	if dormId == "" {
		return &response.GenericError{
			Message: "dormId is missing from query parameters",
			Err:     nil,
		}
	}
	var dorm []model.Dorm

	if result := mysql.Gorm.Where("Id  = ?", dormId).First(&dorm); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get dorm",
			Err:     result.Error,
		}
	}

	return c.JSON(dorm)
}
