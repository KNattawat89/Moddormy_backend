package mange_dorm

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/config"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func GetDormDetail(c *fiber.Ctx) error {
	var dorm model.Dorm

	dormId := c.Query("dormId")
	if dormId == "" {
		return &response.GenericError{
			Message: "dormId is missing from query parameters",
			Err:     nil,
		}
	}

	if result := mysql.Gorm.Where("Id  = ?", dormId).First(&dorm); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get dorm",
			Err:     result.Error,
		}
	}

	coverImage, _ := url.JoinPath(config.C.ProductionURL, *dorm.CoverImage)
	dorm.CoverImage = &coverImage

	return c.JSON(dorm)
}

// var room model.Room

// if result := mysql.Gorm.Where("dorm_id  = ?", dormId).Find(&room); result.Error != nil {
// 	return &response.GenericError{
// 		Message: "Unable to get room",
// 		Err:     result.Error,
// 	}
// }
// body := new(model.Dorm)
// if err := c.BodyParser(body); err != nil {
// 	return &response.GenericError{
// 		Message: "Unable to parse body",
// 		Err:     err,
// 	}
// }

// Dorm := &model.Dorm{

// 	DormName:       body.DormName,
// 	OwnerId:        body.OwnerId,
// 	CoverImage:     body.CoverImage,
// 	HouseNumber:    body.HouseNumber,
// 	Street:         body.Street,
// 	Soi:            body.Soi,
// 	SubDistrict:    body.SubDistrict,
// 	District:       body.District,
// 	City:           body.City,
// 	Zipcode:        body.Zipcode,
// 	Desc:           body.Desc,
// 	AdvancePayment: body.AdvancePayment,
// 	ElectricPrice:  body.ElectricPrice,
// 	WaterPrice:     body.WaterPrice,
// 	Other:          body.Other,
// 	Distant:        body.Distant,
// 	Pet:            body.Pet,
// 	SmokeFree:      body.SmokeFree,
// 	Parking:        body.Parking,
// 	Lift:           body.Lift,
// 	Pool:           body.Pool,
// 	Fitness:        body.Fitness,
// 	Wifi:           body.Wifi,
// 	KeyCard:        body.KeyCard,
// 	CCTV:           body.CCTV,
// 	SecurityGuard:  body.SecurityGuard,
// 	Rooms:          body.Rooms,
// }
