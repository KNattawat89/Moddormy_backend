package upload

import (
	"Moddormy_backend/loaders/mysql"
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/text"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"image"
	"image/jpeg"
	"os"
)

func Rooming(c *fiber.Ctx) error {
	// * Parse user JWT token
	//token := c.Locals("user").(*jwt.Token)
	//claims := token.Claims.(*common.UserClaim)

	// * Parse body
	body := new(payload.UploadRoom)
	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     nil,
		}
	}

	form, _ := c.MultipartForm()
	files := form.File["image"]
	for _, file := range files {
		fileHeader, err := file.Open()
		if err != nil {
			return &response.GenericError{
				Message: "Cannot open image",
				Err:     err,
			}
		}
		defer fileHeader.Close()

		// * Decode image
		img, _, err := image.Decode(fileHeader)
		if err != nil {
			return &response.GenericError{
				Message: "unable to decode file as image",
				Err:     err,
			}
		}

		// * Assign file path
		fileSalt := *text.GenerateString(text.GenerateStringSet.Num, 6)

		// * Save image to file
		savingFile, err := os.Create("./images/" + fileSalt + ".jpeg")
		if err != nil {
			return &response.GenericError{
				Message: "Unable to create an image file",
				Err:     err,
			}
		}
		defer savingFile.Close()

		if err := jpeg.Encode(savingFile, img, nil); err != nil {
			return &response.GenericError{
				Message: "Unable to save an image file",
				Err:     err,
			}
		}

		fileName := fmt.Sprintf("/images/%s.jpeg", fileSalt)
		// * Update user record
		roomImage := &model.RoomImage{
			RoomId:    body.RoomId,
			Room:      nil,
			FileName:  &fileName,
			UpdatedAt: nil,
		}
		if result := mysql.Gorm.Create(roomImage); result.Error != nil {
			return &response.GenericError{
				Message: "Unable to fetch room image",
				Err:     result.Error,
			}
		}
	}

	// * Parse multipart file parameter

	// * Open multipart to file

	return c.JSON(&response.InfoResponse{
		Success: true,
		Message: "Updated image already",
	})
}
