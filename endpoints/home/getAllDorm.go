package home

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/config"
	"Moddormy_backend/utils/value"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"sort"
)

func GetAllDorm(c *fiber.Ctx) error {
	var dorms []model.Dorm
	if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Find(&dorms); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get all dorm",
			Err:     result.Error,
		}
	}

	mappedDorm, _ := value.Iterate(dorms, func(dorm model.Dorm) (*payload.DormSearch, error) {
		//price
		var prices []float64
		for _, room := range dorm.Rooms {
			prices = append(prices, *room.Price)
		}
		sort.Float64sAreSorted(prices)
		//rate
		var overallRates []float64
		for _, rate := range dorm.Reviews {
			overallRates = append(overallRates, *rate.RatingOverall)
		}
		var sum float64
		var finalRate float64
		if len(overallRates) > 0 {
			for i := 0; i < len(overallRates); i++ {
				sum = sum + overallRates[i]
			}
			finalRate = sum / float64(len(overallRates))
		} else {
			finalRate = 0
		}
		//coverimage
		coverImage, _ := url.JoinPath(config.C.ProductionURL, *dorm.CoverImage)

		return &payload.DormSearch{
			DormId:     dorm.Id,
			DormName:   dorm.DormName,
			CoverImage: &coverImage,
			MinPrice:   &prices[0],
			MaxPrice:   &prices[len(prices)-1],
			Rating:     &finalRate,
		}, nil
	})

	return c.JSON(response.NewResponse(mappedDorm))
	//return c.JSON(dorms)
}
