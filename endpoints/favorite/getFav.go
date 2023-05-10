package favorite

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/config"
	"Moddormy_backend/utils/value"
	"fmt"
	"math"
	"net/url"
	"sort"

	"github.com/gofiber/fiber/v2"
)

func GetFav(c *fiber.Ctx) error {
	userId := c.Query("userId")
	if userId == "" {
		return &response.GenericError{
			Message: "userId is missing from query parameters",
			Err:     nil,
		}
	}

	var favs []model.Favorite
	if result := mysql.Gorm.Where("user_id = ?", userId).Find(&favs); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get favorites",
			Err:     result.Error,
		}
	}

	if len(favs) == 0 {
		return &response.GenericError{
			Message: "No favorites found for this user",
			Err:     nil,
		}
	}

	var dorms []model.Dorm
	for _, fav := range favs {
		var dorm model.Dorm
		if result := mysql.Gorm.Preload("Rooms").Preload("Reviews").First(&dorm, fav.DormId); result.Error != nil {
			return &response.GenericError{
				Message: fmt.Sprintf("Unable to get dorm with id %d", fav.DormId),
				Err:     result.Error,
			}
		}
		dorms = append(dorms, dorm)
	}

	mappedDormName, _ := value.Iterate(dorms, func(dorm model.Dorm) (*payload.DormSearch, error) {
		var ratingSum float64
		for _, review := range dorm.Reviews {
			ratingSum += *review.RatingOverall
		}
		ratingAvg := ratingSum / float64(len(dorm.Reviews))
		if math.IsNaN(ratingAvg) {
			ratingAvg = 0
		}

		coverImage, _ := url.JoinPath(config.C.ProductionURL, *dorm.CoverImage)
		var prices []float64
		for _, room := range dorm.Rooms {
			prices = append(prices, *room.Price)
		}
		sort.Float64sAreSorted(prices)
		favStatus := true
		return &payload.DormSearch{
			DormId:      dorm.Id,
			DormName:    dorm.DormName,
			CoverImage:  &coverImage,
			OverallRate: &ratingAvg,
			MinPrice:    &prices[0],
			MaxPrice:    &prices[len(prices)-1],
			FavStatus:   &favStatus,
		}, nil
	})

	return c.JSON(response.NewResponse(mappedDormName))
}

//return c.JSON(dorms) //อันนี้ออกมาเฉพาะfav dorm for each user ยังไม่ได้mapเอาข้อมูลออกมา
//getfav คือใส่query userid ไปด้วยละเอาdormที่มีuseridนั้นๆมาจากfavorite ค่อยมาmapped dormเอาข้อมูลเหมือนในหน้าhome getalldorm
