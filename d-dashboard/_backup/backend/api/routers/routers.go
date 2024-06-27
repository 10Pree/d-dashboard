package routers

import (
	"d-dashboard/controller"

	"github.com/gofiber/fiber/v2"
)

func setupRouters(app *fiber.App) {
	// app.Get("/api", Welcome)
	app.Post("/api/users", controller.CreatedUser)
	app.Get("/api/users", controller.GetUsers)
	app.Get("/api/user/:id", controller.GetUser)
	app.Post("/api/user/:id", controller.UpdateUser)
	app.Delete("/api/user/:id", controller.DeleteUser)
}
