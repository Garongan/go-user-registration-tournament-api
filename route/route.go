package route

import (
	"github.com/gofiber/fiber/v2"
	"go-user-registration-tournament/controller"
	"go-user-registration-tournament/controller/auth"
	"go-user-registration-tournament/controller/config"
	"go-user-registration-tournament/controller/user"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", controller.Hello)

	app.Get("/csrf-token", config.GetCSRFToken)

	users := app.Group("/users")
	users.Post("/sign-up", auth.SignUp)
	users.Post("/sign-in", auth.SignIn)
	users.Get("/", user.GetUser)
}
