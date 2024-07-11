package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-user-registration-tournament/database"
	"go-user-registration-tournament/dto"
	"go-user-registration-tournament/model"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to parse request body",
			Data:       nil,
		})
	}

	var existingAccount model.Account
	if err := database.DB.Where("username = ?", data["username"]).First(&existingAccount).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Account already exists",
			Data:       nil,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to hash password",
			Data:       nil,
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
		return c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to create account",
			Data:       nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.Response{
		StatusCode: fiber.StatusCreated,
		Message:    "Account created successfully",
		Data:       nil,
	})
}
