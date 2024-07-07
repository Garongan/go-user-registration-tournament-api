package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-user-registration-tournament/database"
	"go-user-registration-tournament/route"
	"time"
)

func main() {
	_, err := database.ConnectDB()

	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Connected to database!")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://127.0.0.1:3000",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour / time.Second),
	}))

	route.SetUpRoutes(app)

	err = app.Listen(":8080")
	if err != nil {
		panic("Cannot start server!")
	}
}
