package dto

import "github.com/gofiber/fiber/v2"

func ParseRequest(c *fiber.Ctx) (map[string]string, Response) {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return nil, Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to parse request body",
			Data:       nil,
		}
	}
	return data, Response{}
}
