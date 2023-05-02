package mange_dorm

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func UpdateDorm(c *fiber.Ctx) error {
	dormId := c.Query("dormId")
	var dorm model.Dorm

	// check if dorm with specified ID exists
	if err := mysql.Gorm.Where("Id = ?", dormId).First(&dorm).Error; err != nil {
		return &response.GenericError{
			Message: "Dorm not found",
			Err:     err,
		}
	}

	// update dorm record
	if err := c.BodyParser(&dorm); err != nil {
		return &response.GenericError{
			Message: "Failed to parse request body",
			Err:     err,
		}
	}

	body := new(payload.Dorm)
	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	editDorm := &model.Dorm{

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

	if err := mysql.Gorm.Save(&editDorm).Error; err != nil {
		return &response.GenericError{
			Message: "Failed to update dorm",
			Err:     err,
		}
	}

	return c.JSON(editDorm)
}
