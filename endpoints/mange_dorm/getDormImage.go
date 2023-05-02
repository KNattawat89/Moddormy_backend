package mange_dorm

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func GetDormImage(c *fiber.Ctx) error {
	dormId := c.Query("dormId")
	if dormId == "" {
		return &response.GenericError{
			Message: "dormId is missing from query parameters",
			Err:     nil,
		}
	}

	var dormImage []model.DormImage

	if result := mysql.Gorm.Where("dorm_Id  = ?", dormId).Find(&dormImage); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get dorm",
			Err:     result.Error,
		}
	}

	return c.JSON(dormImage)
}
