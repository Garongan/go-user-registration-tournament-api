package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-user-registration-tournament/database"
	"go-user-registration-tournament/model"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     "Failed to parse request body",
			"data":        nil,
		})
	}

	var existingAccount model.Account
	if err := database.DB.Where("username = ?", data["username"]).First(&existingAccount).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     "Account already exists",
			"data":        nil,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status_code": fiber.StatusInternalServerError,
			"message":     "Failed to hash password",
			"data":        nil,
		})
	}

	accountID := uuid.New().String()

	user := model.User{
		ID:          uuid.New().String(),
		Name:        data["name"],
		PhoneNumber: data["phone"],
		AccountID:   accountID,
	}

	account := model.Account{
		ID:       uuid.New().String(),
		Username: data["username"],
		Password: string(hashedPassword),
		User:     user,
	}

	if err := database.DB.Create(&account).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status_code": fiber.StatusInternalServerError,
			"message":     "Failed to create account",
			"data":        nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status_code": fiber.StatusCreated,
		"message":     "Account created successfully",
		"data":        nil,
	})
}
