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

	// dorm := &model.Dorm{}
	// if result := mysql.Gorm.First(dorm, body.DormId); result.Error != nil {
	// 	return &response.GenericError{
	// 		Message: "Unable to find dorm",
	// 		Err:     result.Error,
	// 	}
	// }

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
		CoverImage:  body.CoverImage,
		Price:       body.Price,
		Desc:        body.Desc,
		Size:        body.Size,
		Airc:        body.RoomFeature.Airc,
		Furniture:   body.RoomFeature.Furniture,
		WaterHeater: body.RoomFeature.WaterHeater,
		Fan:         body.RoomFeature.Fan,
		Fridge:      body.RoomFeature.Fridge,
		Bathroom:    body.RoomFeature.Fridge,
		TV:          body.RoomFeature.TV,
	}

	if result := mysql.Gorm.Create(Room); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to create room",
			Err:     result.Error,
		}
	}

	return c.JSON(Room)
}
