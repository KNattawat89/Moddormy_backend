package mange_room

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func PostRoom(c *fiber.Ctx) error {
	body := new(payload.Room)
	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	dorm := &model.Dorm{}
	if result := mysql.Gorm.First(dorm, body.DormId); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to find dorm",
			Err:     result.Error,
		}
	}

	var room model.Room
	if err := c.BodyParser(&room); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	Room := &model.Room{
		RoomName:    body.RoomName,
		DormId:      body.DormId,
		Dorm:        dorm,
		CoverImage:  body.CoverImage,
		Price:       body.Price,
		Desc:        body.Desc,
		Size:        body.Size,
		Airc:        body.Airc,
		Furniture:   body.Furniture,
		WaterHeater: body.WaterHeater,
		Fan:         body.Fan,
		Fridge:      body.Fridge,
		Bathroom:    body.Fridge,
		TV:          body.TV,
	}

	if result := mysql.Gorm.Create(Room); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to create room",
			Err:     result.Error,
		}
	}

	return c.JSON(Room)
}
