package profile

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/value"

	"github.com/gofiber/fiber/v2"
)

func GetProfileDorm(c *fiber.Ctx) error {
	userId := c.Query("userId")
	if userId == "" {
		return &response.GenericError{
			Message: "userId is missing from query parameters",
			Err:     nil,
		}
	}
	var dorms []model.Dorm

	if result := mysql.Gorm.Where("owner_id = ?", userId).Find(&dorms); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get dorms",
			Err:     result.Error,
		}
	}

	if len(dorms) == 0 {
		return &response.GenericError{
			Message: "No dorms found for this user",
			Err:     nil,
		}
	}

	mappedDormProfile, _ := value.Iterate(dorms, func(dorm model.Dorm) (*payload.DormProfile,error) {
			return &payload.DormProfile{
				DormId: dorm.Id,
				DormName: dorm.DormName,
				CoverImage: dorm.CoverImage,
				CreatedAt: dorm.CreatedAt,
			},nil
	});

	return c.JSON(response.NewResponse(mappedDormProfile))
}
