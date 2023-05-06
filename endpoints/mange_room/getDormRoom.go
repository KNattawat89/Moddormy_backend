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

func GetDormRoom(c *fiber.Ctx) error {

	dormId := c.Query("dormId")
	if dormId == "" {
		return &response.GenericError{
			Message: "dormId is missing from query parameters",
			Err:     nil,
		}
	}

	var room []model.Room

	if result := mysql.Gorm.Where("dorm_id  = ?", dormId).Find(&room); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get room",
			Err:     result.Error,
		}
	}
	roomPayloads := make([]*payload.Room, 0)
	for i := 0; i < len(room); i++ {
		coverImage, _ := url.JoinPath(config.C.ProductionURL, *room[i].CoverImage)
		room[i].CoverImage = &coverImage

		roomPayload := &payload.Room{
			RoomId:     room[i].Id,
			DormId:     room[i].DormId,
			RoomName:   room[i].RoomName,
			CoverImage: room[i].CoverImage,
			Price:      room[i].Price,
			Desc:       room[i].Desc,
			Size:       room[i].Size,
			RoomFeature: &payload.RoomFeature{
				Airc:        room[i].Airc,
				Furniture:   room[i].Furniture,
				WaterHeater: room[i].WaterHeater,
				Fan:         room[i].Fan,
				Fridge:      room[i].Fridge,
				Bathroom:    room[i].Bathroom,
				TV:          room[i].TV,
			},
		}
		roomPayloads = append(roomPayloads, roomPayload)

	}

	return c.JSON(roomPayloads)
}
