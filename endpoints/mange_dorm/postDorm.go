package mange_dorm

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func PostDorm(c *fiber.Ctx) error {

	body := new(payload.Dorm)
	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	var dorm model.Dorm
	if err := c.BodyParser(&dorm); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	Dorm := &model.Dorm{

		DormName:       body.DormName,
		OwnerId:        body.UserId,
		CoverImage:     body.CoverImage,
		HouseNumber:    body.HouseNumber,
		Street:         body.Street,
		Soi:            body.Soi,
		SubDistrict:    body.SubDistrict,
		District:       body.District,
		City:           body.City,
		Zipcode:        body.Zipcode,
		Desc:           body.Desc,
		AdvancePayment: body.AdvancePayment,
		ElectricPrice:  body.ElectricPrice,
		WaterPrice:     body.WaterPrice,
		Other:          body.Other,
		Distant:        body.Distant,
		Pet:            body.DormFeatures.Pet,
		SmokeFree:      body.DormFeatures.SmokeFree,
		Parking:        body.DormFeatures.Parking,
		Lift:           body.DormFeatures.Lift,
		Pool:           body.DormFeatures.Pool,
		Fitness:        body.DormFeatures.Fitness,
		Wifi:           body.DormFeatures.Wifi,
		KeyCard:        body.DormFeatures.KeyCard,
		CCTV:           body.DormFeatures.CCTV,
		SecurityGuard:  body.DormFeatures.SecurityGuard,
	}

	if result := mysql.Gorm.Create(&Dorm); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to create dorm",
			Err:     result.Error,
		}
	}

	return c.JSON(Dorm)
}
