package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-user-registration-tournament/database"
	"go-user-registration-tournament/model"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func CheckPasswordHash(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignIn(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"message":     "Failed to parse request body",
			"data":        nil,
		})
	}

	var account model.Account
	database.DB.Where("username = ?", data["username"]).First(&account)
	if account.ID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status_code": fiber.StatusUnauthorized,
			"message":     "Invalid username or password",
			"data":        nil,
		})
	}

	if CheckPasswordHash([]byte(account.Password), []byte(data["password"])) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status_code": fiber.StatusUnauthorized,
			"message":     "Invalid username or password",
			"data":        nil,
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": account.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status_code": fiber.StatusInternalServerError,
			"message":     "Failed to generate token",
			"data":        nil,
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Success",
		"data": fiber.Map{
			"user_id": user.ID,
			"token":   token,
		},
	})
}
