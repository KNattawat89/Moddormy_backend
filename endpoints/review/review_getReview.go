package review

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func GetDormReview(c *fiber.Ctx) error {
	dormId := c.Query("dormId")
	if dormId == "" {
		return &response.GenericError{
			Message: "dormId is missing from query parameters",
			Err:     nil,
		}
	}

	var review []model.Review

	if result := mysql.Gorm.Where("dorm_id = ?", dormId).Find(&review); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get review",
			Err:     result.Error,
		}
	}

	if len(review) == 0 {
		return &response.GenericError{
			Message: "No reviews found for dormId",
			Err:     nil,
		}
	}

	return c.JSON(review)
}
