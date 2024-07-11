package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"go-user-registration-tournament/config"
	"go-user-registration-tournament/dto"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.Config("JWT_SECRET"))},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Missing ot malformed JWT",
			Data:       nil,
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(dto.Response{
		StatusCode: fiber.StatusUnauthorized,
		Message:    "Invalid or expired JWT",
		Data:       nil,
	})
}
