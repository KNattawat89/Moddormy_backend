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
	// "fmt"
)

func GetDormAll(c *fiber.Ctx) error {
	var dorms []model.Dorm
	if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").Find(&dorms); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get all dorm",
			Err:     result.Error,
		}
	}

	data, _ := value.Iterate(dorms, func(dorm model.Dorm) (*payload.Home, error) {
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
		fav := false

		return &payload.Home{
			DormId:      dorm.Id,
			DormName:    dorm.DormName,
			CoverImage:  &coverImage,
			MinPrice:    &prices[0],
			MaxPrice:    &prices[len(prices)-1],
			OverallRate: &finalRate,
			FavStatus:   &fav,
		}, nil
	})

	// fmt.Println(len(data))
	//find fav
	userId := c.Query("userId")
	if userId == "" {
		return &response.GenericError{
			Message: "userId is missing from query parameters",
			Err:     nil,
		}
	}
	var favDorm []model.Favorite
	if result := mysql.Gorm.Where("user_id = ?", userId).Find(&favDorm); result.Error != nil {
		return &response.GenericError{
			Message: "Fav dorm not found",
			Err:     result.Error,
		}
	}

	// fmt.Println(*data[0].FavStatus)
	// fmt.Println(*favDorm[0].DormId)

	for i := 0; i < len(favDorm); i++ {
		//favStatus
		for j := 0; j < len(data); j++ {
			if *favDorm[i].DormId == *data[j].DormId {
				status := true
				data[j].FavStatus = &status
				break
			}
		}

	}

	return c.JSON(response.NewResponse(data))
	// return c.JSON(favDorm)
}
