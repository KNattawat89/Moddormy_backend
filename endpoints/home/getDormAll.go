package home

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/utils/config"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/value"
	"sort"
	"net/url"
	"github.com/gofiber/fiber/v2"
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
		var overall_rates []float32
		for _, rate := range dorm.Reviews{
			overall_rates = append(overall_rates, *rate.RatingOverall)
		}
		var sum float32
		var final_rate float32
		if len(overall_rates) >0 {
			for i:=0; i < len(overall_rates);i++{
				sum = sum + overall_rates[i]
			}
			final_rate = sum/float32(len(overall_rates))
		} else {
			final_rate =0 
		}
		//coverimage
		coverImage, _ := url.JoinPath(config.C.URL, *dorm.CoverImage)
		fav := false

		return &payload.Home{
			DormId: dorm.Id,
			DormName: dorm.DormName,
			CoverImage: &coverImage,
			MinPrice:   &prices[0],
			MaxPrice:   &prices[len(prices)-1],
			OverallRate: &final_rate,
			FavStatus: &fav,
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
	var fav_dorm []model.Favorite
	if result := mysql.Gorm.Where("user_id = ?", userId).Find(&fav_dorm); result.Error != nil {
		return &response.GenericError{
			Message: "Fav dorm not found",
			Err:     result.Error,
		}
	}
	
	// fmt.Println(*data[0].FavStatus)
	// fmt.Println(*fav_dorm[0].DormId)

	for i:= 0 ; i < len(fav_dorm);i++{
		//favStatus
		for j := 0; j < len(data); j++ {
			if *fav_dorm[i].DormId == *data[j].DormId {
				status := true
				data[j].FavStatus = &status
				break
			}
		}
		
	}
	
	return c.JSON(response.NewResponse(data))
	// return c.JSON(fav_dorm)
}