package middlewares

import (
	"Moddormy_backend/types/common"
	"errors"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Jwt() fiber.Handler {
	conf := jwtware.Config{
		SigningKey:  []byte("123456789"),
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		ContextKey:  "user",
		Claims:      &common.UserClaims{},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return errors.New("Unauthorized")
		},
	}

	return jwtware.New(conf)
}
