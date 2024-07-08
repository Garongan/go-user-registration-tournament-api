package route

import (
	"github.com/gofiber/fiber/v2"
	"go-user-registration-tournament/controller"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", controller.Hello)
	app.Get("/csrf-token", controller.GetCSRFToken)
	app.Post("users/sign-up", controller.SignUp)
}
