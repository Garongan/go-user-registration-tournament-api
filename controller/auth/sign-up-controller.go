package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-user-registration-tournament/database"
	"go-user-registration-tournament/dto"
	"go-user-registration-tournament/middleware"
	"go-user-registration-tournament/model"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	request, response := dto.ParseRequest(c)
	if response != (dto.Response{}) {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response = middleware.SignUpValidation(request["name"], request["phone"], request["username"], request["password"])
	if response != (dto.Response{}) {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	var existingAccount model.Account
	if err := database.DB.Where("username = ?", request["username"]).First(&existingAccount).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Account already exists",
			Data:       nil,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request["password"]), bcrypt.DefaultCost)
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
		Name:        request["name"],
		PhoneNumber: request["phone"],
		AccountID:   accountID,
	}

	account := model.Account{
		ID:       uuid.New().String(),
		Username: request["username"],
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
