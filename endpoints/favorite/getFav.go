package favorite

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"

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
	//var dorms []model.Dorm
	var favs []model.Favorite
	if result := mysql.Gorm.Where("user_id = ?", userId).Find(&favs); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get dorms",
			Err:     result.Error,
		}
	}

	if len(favs) == 0 {
		return &response.GenericError{
			Message: "No dorms found for this user's favorite",
			Err:     nil,
		}
	}

	// mappedDormName, _ := value.Iterate(favs, func(fav model.Favorite) (*payload.DormSearch, error) {
	// 	var ratingSum float32
    //     for _, review := range dorm.Reviews {
    //         ratingSum += *review.RatingOverall
    //     }
    //     ratingAvg := ratingSum / float32(len(dorm.Reviews))

    //     coverImage, _ := url.JoinPath(config.C.URL, *dorm.CoverImage)
    //     var prices []float64
    //     for _, room := range dorm.Rooms {
    //         prices = append(prices, *room.Price)
    //     }
    //     sort.Float64sAreSorted(prices)

    //     return &payload.DormSearch{
    //         DormId:     dorm.Id,
    //         DormName:   dorm.DormName,
    //         CoverImage: &coverImage,
    //         Rating:     &ratingAvg,
    //         MinPrice:   &prices[0],
    //         MaxPrice:   &prices[len(prices)-1],
    //     }, nil
	// })

	//return c.JSON(response.NewResponse(mappedDormName))
	return c.JSON(favs) //อันนี้ออกมาเฉพาะfav dorm for each user ยังไม่ได้mapเอาข้อมูลออกมา
	//getfav คือใส่query userid ไปด้วยละเอาdormที่มีuseridนั้นๆมาจากfavorite ค่อยมาmapped dormเอาข้อมูลเหมือนในหน้าhome getalldorm
}
