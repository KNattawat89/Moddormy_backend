package review

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"

	"github.com/gofiber/fiber/v2"
	//"github.com/golang-jwt/jwt/v4"
)

func AddDormReview(c *fiber.Ctx) error {
	// * Parse user
	//u := c.Locals("user").(*jwt.Token).Claims.(*common.UserClaims)

	body := new(payload.DormReview)
	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	// Get the user and dorm models based on the given IDs

	user := &model.User{}
	if result := mysql.Gorm.First(user, body.UserId); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to find user",
			Err:     result.Error,
		}
	}

	dorm := &model.Dorm{}
	if result := mysql.Gorm.First(dorm, body.DormId); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to find dorm",
			Err:     result.Error,
		}
	}

	//แก้ชั่วคราว เพราะเปลี่ยน overall -> float
	overall := float64(*body.RatingOverall)

	DormReview := &model.Review{
		UserId: body.UserId,

		DormId: body.DormId,

		Review:         body.Review,
		RatingPrice:    body.RatingPrice,
		RatingLocation: body.RatingLocation,
		RatingFacility: body.RatingFacility,
		RatingSanitary: body.RatingSanitary,
		RatingSecurity: body.RatingSecurity,
		RatingOverall:  &overall,
		CreatedAt:      nil,
	}

	if result := mysql.Gorm.Create(&DormReview); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to create review",
			Err:     result.Error,
		}
	}

	return c.JSON(DormReview)
}
