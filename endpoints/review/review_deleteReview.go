package review

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func DeleteDormReview(c *fiber.Ctx) error {

	dormId := c.Query("dormId")
	userId := c.Query("userId")

	if dormId == "" {
		return &response.GenericError{
			Message: "dormId is missing from query parameters",
			Err:     nil,
		}
	}

	if userId == "" {
		return &response.GenericError{
			Message: "userId is missing from query parameters",
			Err:     nil,
		}
	}

	// Check if the review exists
	var review model.Review
	if result := mysql.Gorm.Where("user_id = ? AND dorm_id = ?", userId, dormId).First(&review); result.Error != nil {
		return &response.GenericError{
			Message: "Review not found",
			Err:     result.Error,
		}
	}

	// Delete the review
	if result := mysql.Gorm.Delete(&review); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to delete review",
			Err:     result.Error,
		}
	}

	return c.JSON(&response.GenericSuccess{
		Message: "Review deleted successfully",
	})

}
