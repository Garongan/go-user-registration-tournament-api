package route

import (
	"github.com/gofiber/fiber/v2"
	"go-user-registration-tournament/controller"
	"go-user-registration-tournament/controller/auth"
	"go-user-registration-tournament/controller/config"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", controller.Hello)

	app.Get("/csrf-token", config.GetCSRFToken)

	app.Post("/users/sign-up", auth.SignUp)
	app.Post("/users/sign-in", auth.SignIn)
}
