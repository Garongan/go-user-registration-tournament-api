package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-user-registration-tournament/database"
	"go-user-registration-tournament/model"
	"os"
)

func GetUser(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims, ok := token.Claims.(*jwt.MapClaims)

	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to parse claims",
		})
	}

	accountID, _ := (*claims)["sub"].(string)
	account := model.Account{ID: accountID}

	database.DB.Where("id = ?", accountID).First(&account)

	var user model.User
	database.DB.Where("account_id = ?", account.ID).First(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":       user.ID,
		"name":     user.Name,
		"phone":    user.PhoneNumber,
		"username": account.Username,
	})
}
