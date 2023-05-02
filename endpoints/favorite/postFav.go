package favorite

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PostFav (c *fiber.Ctx) error {

    body := new(payload.Favorite)
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

    dorm := &model.Dorm{}
	if result := mysql.Gorm.First(dorm, body.DormId); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to find dorm",
			Err:     result.Error,
		}
	}

    now := time.Now()

    // Format the date and time as a string in the required format for MySQL
    updatedAt := now.Format("2006-01-02 15:04:05")

    // Parse the createAt string into a time.Time value
    updatedAtTime, err := time.Parse("2006-01-02 15:04:05", updatedAt)
    if err != nil {
        log.Println("handle error")
    }

    Favorite := &model.Favorite{

        DormId:       body.DormId,
        UserId:        body.UserId,
        UpdatedAt:        &updatedAtTime,
        IsFav:         body.IsFav,
    }

    if result := mysql.Gorm.Create(&Favorite); result.Error != nil {
        return &response.GenericError{
            Message: "Unable to favorite dorm",
            Err:     result.Error,
        }
    }

    return c.JSON(Favorite)
}