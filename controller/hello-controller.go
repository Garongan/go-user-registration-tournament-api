package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-user-registration-tournament/dto"
)

func Hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(dto.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Success from Hello Controller",
		Data:       "Hello, From Go User Registration Tournament API! 😊🚀💖",
	})
}
