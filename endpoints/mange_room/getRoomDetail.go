package mange_room

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/config"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func GetRoomDetail(c *fiber.Ctx) error {
	roomId := c.Query("roomId")
	if roomId == "" {
		return &response.GenericError{
			Message: "roomId is missing from query parameters",
			Err:     nil,
		}
	}
	var room model.Room

	if result := mysql.Gorm.Where("Id  = ?", roomId).First(&room); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get room",
			Err:     result.Error,
		}
	}

	coverImage, _ := url.JoinPath(config.C.ProductionURL, *room.CoverImage)
	room.CoverImage = &coverImage

	roomPayload := &payload.Room{
		RoomId:     room.Id,
		DormId:     room.DormId,
		RoomName:   room.RoomName,
		CoverImage: room.CoverImage,
		Price:      room.Price,
		Desc:       room.Desc,
		Size:       room.Size,
		RoomFeature: &payload.RoomFeature{
			Airc:        room.Airc,
			Furniture:   room.Furniture,
			WaterHeater: room.WaterHeater,
			Fan:         room.Fan,
			Fridge:      room.Fridge,
			Bathroom:    room.Bathroom,
			TV:          room.TV,
		},
	}

	return c.JSON(roomPayload)
}
