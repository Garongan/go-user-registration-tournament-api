package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"go-user-registration-tournament/config"
	"go-user-registration-tournament/dto"
	"regexp"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.Config("JWT_SECRET"))},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(convertToResponse("Missing ot malformed JWT"))
	}
	return c.Status(fiber.StatusUnauthorized).JSON(dto.Response{
		StatusCode: fiber.StatusUnauthorized,
		Message:    "Invalid or expired JWT",
		Data:       nil,
	})
}

func convertToResponse(message string) dto.Response {
	return dto.Response{
		StatusCode: fiber.StatusBadRequest,
		Message:    message,
		Data:       nil,
	}
}

func SignInValidation(username string, password string) dto.Response {
	ok := len(username) >= 6 && len(username) <= 15
	if !ok {
		return convertToResponse("Username length must be between 6 and 15 characters")
	}

	pattern := regexp.MustCompile("^[a-zA-Z0-9]+$")
	ok = pattern.MatchString(username)
	if !ok {
		return convertToResponse("Username should only use alphabet and numbers")
	}

	ok = len(password) >= 8 && len(password) <= 20
	if !ok {
		return convertToResponse("Password length must be between 8 and 20 characters")
	}

	return dto.Response{}
}

func SignUpValidation(name, phone, username, password string) dto.Response {
	ok := len(name) >= 3 && len(name) <= 50
	if !ok {
		return convertToResponse("Name length must be between 3 and 50 characters")
	}

	pattern := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	ok = pattern.MatchString(name)
	if !ok {
		return convertToResponse("Name should only use alphabet and space")
	}

	ok = len(phone) >= 7 && len(phone) <= 14
	if !ok {
		return convertToResponse("Phone number length must be between 10 and 15 characters")
	}

	pattern = regexp.MustCompile(`^\+[0-9]\d{1,14}$`)
	ok = pattern.MatchString(phone)
	if !ok {
		return convertToResponse("Phone number should be valid")
	}

	ok = len(username) >= 6 && len(username) <= 15
	if !ok {
		return convertToResponse("Username length must be between 6 and 15 characters")
	}

	pattern = regexp.MustCompile("^[a-zA-Z0-9]+$")
	ok = pattern.MatchString(username)
	if !ok {
		return convertToResponse("Username should only use alphabet and numbers")
	}

	ok = len(password) >= 8 && len(password) <= 20
	if !ok {
		return convertToResponse("Password length must be between 8 and 20 characters")
	}

	return dto.Response{}
}
