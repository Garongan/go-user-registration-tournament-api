package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-user-registration-tournament/database"
	"go-user-registration-tournament/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	var existingAccount models.Account
	if err := database.DB.Where("username = ?", data["username"]).First(&existingAccount).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Account already exists",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	accountID := uuid.New().String()

	user := models.User{
		ID:          uuid.New().String(),
		Name:        data["name"],
		PhoneNumber: data["phone"],
		AccountID:   accountID,
	}

	account := models.Account{
		ID:       uuid.New().String(),
		Username: data["username"],
		Password: string(hashedPassword),
		User:     user,
	}

	if err := database.DB.Create(&account).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create account",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Account created successfully",
	})
}
