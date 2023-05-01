package mange_dorm

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/value"
	"github.com/gofiber/fiber/v2"
)

func GetAllDorm(c *fiber.Ctx) error {
	var dorms []model.Dorm
	if result := mysql.Gorm.Select("dorm_name").Find(&dorms); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get all dorm",
			Err:     result.Error,
		}
	}

	mappedDormName, _ := value.Iterate(dorms, func(dorm model.Dorm) (*payload.DormSearch, error) {
		return &payload.DormSearch{
			DormName: dorm.DormName,
		}, nil
	})

	return c.JSON(response.NewResponse(mappedDormName))
}
