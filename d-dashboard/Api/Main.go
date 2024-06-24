package main

import (
	"d-dashboard/connect"
	"d-dashboard/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome")
}

func setupRooutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	app.Get("/api", Welcome)
	app.Post("/api/users", controller.CreatedUser)
	app.Get("/api/users", controller.GetUsers)
	app.Get("/api/user/:id", controller.GetUser)
	app.Post("/api/user/:id", controller.UpdateUser)
	app.Delete("/api/user/:id", controller.DeleteUser)
}

func main() {
	connect.ConnectDb()

	app := fiber.New()
	setupRooutes(app)
	app.Listen(":8080")
}
