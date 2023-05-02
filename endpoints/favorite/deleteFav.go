package favorite

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func DeleteFav(c *fiber.Ctx) error {

	userId := c.Query("userId")
	dormId := c.Query("dormId")

	if userId == "" {
		return &response.GenericError{
			Message: "userId is missing from query parameters",
			Err:     nil,
		}
	}

	if dormId == "" {
		return &response.GenericError{
			Message: "dormId is missing from query parameters",
			Err:     nil,
		}
	}

	// Check if the item exists
	var fav model.Favorite
	if result := mysql.Gorm.Where("user_id = ? AND dorm_id = ?", userId, dormId).First(&fav); result.Error != nil {
		return &response.GenericError{
			Message: "Favorite not found",
			Err:     result.Error,
		}
	}

	// Delete 
	if result := mysql.Gorm.Delete(&fav); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to delete favorite",
			Err:     result.Error,
		}
	}

	return c.JSON(&response.GenericSuccess{
		Message: "Favorite deleted successfully",
	})

}
