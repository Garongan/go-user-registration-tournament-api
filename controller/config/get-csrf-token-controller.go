package config

import (
	"github.com/gofiber/fiber/v2"
)

func GetCSRFToken(c *fiber.Ctx) error {
	token := c.Locals("csrf")
	if token == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get CSRF token",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"csrf_token": token})
}
