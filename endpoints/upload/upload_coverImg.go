package upload

import (
	"Moddormy_backend/types/payload"
	"Moddormy_backend/types/response"
	"Moddormy_backend/utils/text"
	"image"
	"image/jpeg"
	"os"

	"github.com/gofiber/fiber/v2"
)

func CoverImg(c *fiber.Ctx) error {

	fileHeader, err := c.FormFile("image")
	if err != nil {
		return &response.GenericError{
			Message: "Unable to parse image file",
			Err:     err,
		}
	}
	file, err := fileHeader.Open()
	if err != nil {
		return &response.GenericError{
			Message: "Unable to parse image file",
			Err:     err,
		}
	}

	// * Decode image
	img, _, err := image.Decode(file)
	if err != nil {
		return &response.GenericError{
			Message: "Unable to decode image",
			Err:     err,
		}
	}

	fileSalt := *text.GenerateString(text.GenerateStringSet.Num, 6)

	// * Save image to file
	savingFile, err := os.Create("/images/" + fileSalt + ".jpeg")
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

	var image = "./images/" + fileSalt + ".jpeg"

	return c.JSON(
		payload.UploadCoverImg{Image: &image},
	)

}
