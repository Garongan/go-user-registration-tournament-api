package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-user-registration-tournament/database"
	"go-user-registration-tournament/dto"
	"go-user-registration-tournament/middleware"
	"go-user-registration-tournament/model"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func CheckPasswordHash(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

func SignIn(c *fiber.Ctx) error {
	request, response := dto.ParseRequest(c)
	if response != (dto.Response{}) {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response = middleware.SignInValidation(request["username"], request["password"])
	if response != (dto.Response{}) {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	var account model.Account
	database.DB.Where("username = ?", request["username"]).First(&account)
	if account.ID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.Response{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Invalid username or password",
			Data:       nil,
		})
	}

	if CheckPasswordHash([]byte(account.Password), []byte(request["password"])) {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.Response{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Invalid username or password",
			Data:       nil,
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": account.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to generate token",
			Data:       nil,
		})
	}

	cookie := fiber.Cookie{
		Name:        "jwt",
		Value:       token,
		Expires:     time.Now().Add(time.Hour * 24),
		Secure:      true,
		HTTPOnly:    true,
		SessionOnly: true,
	}

	c.Cookie(&cookie)

	var user model.User
	database.DB.Where("account_id = ?", account.ID).First(&user)

	dataResponse := fiber.Map{
		"user_id": user.ID,
		"token":   token,
	}
	return c.Status(fiber.StatusOK).JSON(dto.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Login Success",
		Data:       dataResponse,
	})
}
