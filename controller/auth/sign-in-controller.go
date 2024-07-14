package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-user-registration-tournament/database"
	"go-user-registration-tournament/dto"
	"go-user-registration-tournament/model"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
	"time"
)

func CheckPasswordHash(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignIn(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to parse request body",
			Data:       nil,
		})
	}

	ok := len(data["username"]) >= 6 && len(data["username"]) <= 15
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Username length must be between 6 and 15 characters",
			Data:       nil,
		})
	}

	pattern := regexp.MustCompile("^[a-zA-Z0-9]+$")
	ok = pattern.MatchString(data["username"])
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Username should only use alphabet and numbers",
			Data:       nil,
		})
	}

	ok = len(data["password"]) >= 8 && len(data["password"]) <= 20
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Password length must be between 8 and 20 characters",
			Data:       nil,
		})
	}

	var account model.Account
	database.DB.Where("username = ?", data["username"]).First(&account)
	if account.ID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.Response{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Invalid username or password",
			Data:       nil,
		})
	}

	if CheckPasswordHash([]byte(account.Password), []byte(data["password"])) {
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
