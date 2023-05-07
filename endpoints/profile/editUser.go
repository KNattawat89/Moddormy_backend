package profile

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"

	//"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	//"log"
	//"time"

	"github.com/gofiber/fiber/v2"
)

func EditUser(c *fiber.Ctx) error {
    var user model.User

	// Get user ID from query parameter
	userId := c.Query("userId")
	if userId == "" {
		return &response.GenericError{
			Message: "userId is missing from query parameters",
			Err:     nil,
		}
	}

	// Find the user in the database
	if result := mysql.Gorm.Where("id = ?", userId).First(&user); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to find user in database",
			Err:     result.Error,
		}
	}

	// Update user fields with request body
	if err := c.BodyParser(&user); err != nil {
		return &response.GenericError{
			Message: "Unable to parse request body",
			Err:     err,
		}
	}

	// Save the updated user to the database
	if result := mysql.Gorm.Save(&user); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to update user",
			Err:     result.Error,
		}
	}

	// Return the updated user in JSON format
	return c.JSON(user)
}

