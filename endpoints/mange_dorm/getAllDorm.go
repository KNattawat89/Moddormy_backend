package mange_dorm

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/config"
	"Moddormy_backend/utils/value"
	"math/rand"
	"net/url"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllDorm(c *fiber.Ctx) error {
	var dorms []model.Dorm
	if result := mysql.Gorm.Preload("Rooms").Find(&dorms); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get all dorm",
			Err:     result.Error,
		}
	}

	mappedDormName, _ := value.Iterate(dorms, func(dorm model.Dorm) (*payload.DormSearch, error) {
		//random 1-5 rate เดี๋ยวเปลี่ยน
		rand.Seed(time.Now().UnixNano())
		rating := rand.Intn(5) + 1
		rating2 := float32(rating)
		coverImage, _ := url.JoinPath(config.C.URL, *dorm.CoverImage)
		var prices []float64
		for _, room := range dorm.Rooms {
			prices = append(prices, *room.Price)
		}
		sort.Float64sAreSorted(prices)

		//fmt.Println(prices)
		return &payload.DormSearch{
			DormId:     dorm.Id,
			DormName:   dorm.DormName,
			CoverImage: &coverImage,
			Rating:     &rating2,
			MinPrice:   &prices[0],
			MaxPrice:   &prices[len(prices)-1],
		}, nil
	})

	return c.JSON(response.NewResponse(mappedDormName))
	//return c.JSON(dorms)
}
