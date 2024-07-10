package config

import (
	"github.com/gofiber/fiber/v2"
)

func GetCSRFToken(c *fiber.Ctx) error {
	token := c.Locals("csrf")
	if token == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status_code": fiber.StatusInternalServerError,
			"message":     "Failed to get CSRF token",
			"data":        nil,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status_code": fiber.StatusCreated,
		"message":     "CSRF token generated successfully",
		"data": fiber.Map{
			"csrf_token": token,
		},
	})
}
