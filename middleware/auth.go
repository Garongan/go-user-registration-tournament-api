package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"go-user-registration-tournament/config"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.Config("JWT_SECRET"))},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"status_code": fiber.StatusBadRequest,
				"message":     "Missing or malformed JWT",
				"data":        nil,
			})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{
			"status_code": fiber.StatusUnauthorized,
			"message":     "Invalid or expired JWT",
			"data":        nil,
		})
}
