package profile

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	userId := c.Query("userId")
	if userId == "" {
		return &response.GenericError{
			Message: "userId is missing from query parameters",
			Err:     nil,
		}
	}
	var user []model.User

	if result := mysql.Gorm.Where("Id  = ?", userId).First(&user); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to find user",
			Err:     result.Error,
		}
	}

	return c.JSON(user)
}
