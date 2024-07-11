package router

import (
	"github.com/gofiber/fiber/v2"
	"go-user-registration-tournament/controller"
	"go-user-registration-tournament/controller/auth"
	"go-user-registration-tournament/controller/config"
	"go-user-registration-tournament/controller/user"
	"go-user-registration-tournament/middleware"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", controller.Hello)

	app.Get("/csrf-token", config.GetCSRFToken)

	app.Post("/sign-up", auth.SignUp)
	app.Post("/sign-in", auth.SignIn)

	users := app.Group("/users", middleware.Protected())
	users.Get("/:id", user.GetUser)
}
