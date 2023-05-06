package mange_dorm

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
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

	// dormPayload := payload.Dorm{
	// 	DormId:   dorm.id,
	// 	DormName: dorm.dorm_name,
	// 	CoverImage:     dorm.cover_image,
	// 	HouseNumber:    dorm.house_number,
	// 	Street:       dorm.street,
	// 	Soi: 		dorm.soi,
	// 	SubDistrict:  dorm.sub_district,
	// 	District:    dorm.district,
	// 	City:       		dorm.city,
	// 	Zipcode:     dorm.zipcode,
	// 	Desc:          dorm.desc,
	// 	AdvancePayment: dorm.advance_payment,
	// 	ElectricPrice:  dorm.electric_price,
	// 	WaterPrice:     dorm.water_price,
	// 	Other:         dorm.Other,
	// 	Distant:        dorm.distant,
	// 	DormFeatures:   &payload.DormFeatures{
	// 		Pet:       dorm.pet,
	// 		SmokeFree: dorm.smoke_free,
	// 		Parking:   dorm.parking,
	// 		Lift:      dorm.lift,
	// 		Pool:      dorm.pool,
	// 		Fitness:   dorm.fitness,
	// 		Wifi:      dorm.wifi,
	// 		KeyCard:   dorm.key_card,
	// 		CCTV:      dorm.cctv,
	// 		SecurityGuard: dorm.security_guard,
	// 	},
	// }

	if result := mysql.Gorm.Where("Id  = ?", dormId).First(&dorm); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get dorm",
			Err:     result.Error,
		}
	}

	coverImage, _ := url.JoinPath(config.C.ProductionURL, *dorm.CoverImage)
	dorm.CoverImage = &coverImage
	dormPayload := payload.Dorm{
		DormId:         dorm.Id,
		UserId:         dorm.OwnerId,
		DormName:       dorm.DormName,
		CoverImage:     dorm.CoverImage,
		HouseNumber:    dorm.HouseNumber,
		Street:         dorm.Street,
		Soi:            dorm.Soi,
		SubDistrict:    dorm.SubDistrict,
		District:       dorm.District,
		City:           dorm.City,
		Zipcode:        dorm.Zipcode,
		Desc:           dorm.Desc,
		AdvancePayment: dorm.AdvancePayment,
		ElectricPrice:  dorm.ElectricPrice,
		WaterPrice:     dorm.WaterPrice,
		Other:          dorm.Other,
		Distant:        dorm.Distant,
		DormFeatures: &payload.DormFeature{
			Pet:           dorm.Pet,
			SmokeFree:     dorm.SmokeFree,
			Parking:       dorm.Parking,
			Lift:          dorm.Lift,
			Pool:          dorm.Pool,
			Fitness:       dorm.Fitness,
			Wifi:          dorm.Wifi,
			KeyCard:       dorm.KeyCard,
			CCTV:          dorm.CCTV,
			SecurityGuard: dorm.SecurityGuard,
		},
	}

	return c.JSON(dormPayload)
}
