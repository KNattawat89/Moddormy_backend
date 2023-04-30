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

	user := &model.User{}
	if result := mysql.Gorm.First(user, body.UserId); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to find user",
			Err:     result.Error,
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
		Pet:            body.Pet,
		SmokeFree:      body.SmokeFree,
		Parking:        body.Parking,
		Lift:           body.Lift,
		Pool:           body.Pool,
		Fitness:        body.Fitness,
		Wifi:           body.Wifi,
		KeyCard:        body.KeyCard,
		CCTV:           body.CCTV,
		SecurityGuard:  body.SecurityGuard,
	}

	if result := mysql.Gorm.Create(&Dorm); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to create dorm",
			Err:     result.Error,
		}
	}

	return c.JSON(Dorm)
}
