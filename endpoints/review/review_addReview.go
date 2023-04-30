package review

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"

	"time"

	"github.com/gofiber/fiber/v2"
)

func AddDormReview(c *fiber.Ctx) error {

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

	//t := time.Now()
	//createdAt, _ := time.Parse(time.RFC3339, "2023-04-26")
	//createdAt := time.Now().Format(time.RFC3339)
	//createdAt, err := time.Parse("2006-01-02 15:04:05", "2023-04-26 12:34:56")
	// if err != nil {
	// 	// handle error
	// 	print("error arai wa")
	// }
	// Create the review model
	//createAtTime, err := time.Parse("2006-01-02 15:04:05", createdAt)
	// if err != nil {
	// 	print("error arai wa")
	// }

	now := time.Now()

	// Format the date and time as a string in the required format for MySQL
	createAt := now.Format("2006-01-02 15:04:05")

	// Parse the createAt string into a time.Time value
	createAtTime, err := time.Parse("2006-01-02 15:04:05", createAt)
	if err != nil {
		// handle error
	}
	DormReview := &model.Review{
		UserId:         body.UserId,
		User:           user,
		DormId:         body.DormId,
		Dorm:           dorm,
		Review:         body.Review,
		RatingPrice:    body.RatingPrice,
		RatingLocation: body.RatingLocation,
		RatingFacility: body.RatingFacility,
		RatingSanitary: body.RatingSanitary,
		RatingSecurity: body.RatingSecurity,
		RatingOverall:  body.RatingOverall,
		CreateAt:       &createAtTime,
	}

	if result := mysql.Gorm.Create(&DormReview); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to create review",
			Err:     result.Error,
		}
	}

	return c.JSON(DormReview)
}
