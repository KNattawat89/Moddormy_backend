package fiber

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"Moddormy_backend/types"
)

func defaultErrorHandler(ctx *fiber.Ctx, err error) error {

	// Case of *fiber.Error.
	if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).JSON(types.RespError{
			Success: false,
			Error:   strings.ReplaceAll(strings.ToUpper(e.Message), " ", "_"),
		})
	}

	// Case of *types.PassError, the generic pass-trough error type.
	if e, ok := err.(*types.PassError); ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.RespError{
			Success: false,
			Error:   e.Message,
			Params:  e.Params,
		})
	}

	// Case of validator.ValidationErrors
	if e, ok := err.(validator.ValidationErrors); ok {
		// Construct invalid validation parameters
		var params []string
		for _, err := range e {

			params = append(params, strings.ToLower(err.Field())+","+err.Tag())
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(types.RespError{
			Success: false,
			Error:   "INVALID_VALIDATION",
			Params:  params,
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(
		types.RespError{
			Success: false,
			Error:   "UNKNOWN_SERVER_SIDE_ERROR",
			Params:  []string{err.Error()},
		},
	)
}
