package mange_room

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/config"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func GetRoomImage(c *fiber.Ctx) error {
	roomId := c.Query("roomId")
	if roomId == "" {
		return &response.GenericError{
			Message: "roomId is missing from query parameters",
			Err:     nil,
		}
	}

	var roomImage []model.RoomImage

	if result := mysql.Gorm.Where("room_Id  = ?", roomId).Find(&roomImage); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get room",
			Err:     result.Error,
		}
	}
	for i := 0; i < len(roomImage); i++ {
		images, _ := url.JoinPath(config.C.ProductionURL, *roomImage[i].FileName)
		roomImage[i].FileName = &images
	}

	return c.JSON(roomImage)
}
