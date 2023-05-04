package authentication

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	account := new(payload.RegisterAccount)
	if err := c.BodyParser(account); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     err,
		}
	}
	user := &model.User{
		Id:           nil,
		ProfileImage: nil,
		UserName:     account.UserName,
		Fname:        account.FirstName,
		Lname:        account.LastName,
		Email:        account.Email,
		Tel:          account.Tel,
		LineID:       nil,
		UserType:     account.Account,
		UpdatedAt:    nil,
		UnusedField:  "",
	}
	if result := mysql.Gorm.Create(user); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to create an account",
			Err:     result.Error,
		}
	}
	return c.JSON(&response.InfoResponse{
		Success: true,
		Message: "Created an account already",
	})
}
