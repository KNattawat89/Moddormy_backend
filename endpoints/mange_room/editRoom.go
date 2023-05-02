package mange_room

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func UpdateRoom(c *fiber.Ctx) error {
	roomId := c.Query("roomId")
	var room model.Room

	// check if dorm with specified ID exists
	if err := mysql.Gorm.Where("Id = ?", roomId).First(&room).Error; err != nil {
		return &response.GenericError{
			Message: "Room not found",
			Err:     err,
		}
	}

	// update dorm record
	if err := c.BodyParser(&room); err != nil {
		return &response.GenericError{
			Message: "Failed to parse request body",
			Err:     err,
		}
	}

	body := new(payload.Room)
	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	editRoom := &model.Room{
		RoomName:    body.RoomName,
		DormId:      body.DormId,
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

	if err := mysql.Gorm.Save(&editRoom).Error; err != nil {
		return &response.GenericError{
			Message: "Failed to update room",
			Err:     err,
		}
	}

	return c.JSON(editRoom)
}
