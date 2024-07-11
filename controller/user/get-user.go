package user

import (
	"github.com/gofiber/fiber/v2"
	"go-user-registration-tournament/database"
	"go-user-registration-tournament/dto"
	"go-user-registration-tournament/model"
)

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User
	db.Where("id = ?", id).First(&user)
	if user.ID == "" {
		return c.Status(fiber.StatusNotFound).JSON(dto.Response{
			StatusCode: fiber.StatusNotFound,
			Message:    "No user found with ID",
			Data:       nil,
		})
	}
	var account model.Account
	db.Where("id = ?", user.AccountID).First(&account)
	if account.ID == "" {
		return c.Status(fiber.StatusNotFound).JSON(dto.Response{
			StatusCode: fiber.StatusNotFound,
			Message:    "No account found with ID",
			Data:       nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":       user.ID,
		"name":     user.Name,
		"phone":    user.PhoneNumber,
		"username": account.Username,
	})
}
