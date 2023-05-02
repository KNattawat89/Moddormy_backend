package mange_dorm

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func UpdateDorm(c *fiber.Ctx) error {
	var dorm model.Dorm

	// Get dorm ID from query parameter
	dormId := c.Query("dormId")
	if dormId == "" {
		return &response.GenericError{
			Message: "dormId is missing from query parameters",
			Err:     nil,
		}
	}

	// Find the dorm in the database
	if result := mysql.Gorm.Where("id = ?", dormId).First(&dorm); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to find dorm",
			Err:     result.Error,
		}
	}

	// Update dorm fields with request body
	if err := c.BodyParser(&dorm); err != nil {
		return &response.GenericError{
			Message: "Unable to parse request body",
			Err:     err,
		}
	}

	// Save the updated dorm to the database
	if result := mysql.Gorm.Save(&dorm); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to update dorm",
			Err:     result.Error,
		}
	}

	// Return the updated dorm in JSON format
	return c.JSON(dorm)
}
