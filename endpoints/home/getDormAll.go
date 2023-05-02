package home

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/value"
	"sort"

	"github.com/gofiber/fiber/v2"
)

func GetAllDorm(c *fiber.Ctx) error {
	var dorms []model.Dorm
	if result := mysql.Gorm.Preload("Rooms").Preload("Review").Find(&dorms); result.Error != nil {
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
		var overall_rates []int
		for _, rate := range dorm.Reviews{
			overall_rates = append(overall_rates, *rate.RatingOverall)
		}
		var sum float64
		for i:=0; i < len(overall_rates);i++{
			sum = sum + float64(overall_rates[i])
		}
		var final_rate  := sum/len(overall_rates)

		return &payload.Home{
			DormId: dorm.Id,
			DormName: dorm.DormName,
			CoverImage: dorm.CoverImage,
			MinPrice:   &prices[0],
			MaxPrice:   &prices[len(prices)-1],
			OverallRate: final_rate,
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

	for i:= 0 ; i < len(fav_dorm);i++{
		//favStatus
		
	}
	
	// return c.JSON(response.NewResponse(data))
	return c.JSON(fav_dorm)
}
