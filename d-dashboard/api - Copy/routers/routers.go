package routers

import (
	"d-dashboard/controller"

	"github.com/gofiber/fiber/v2"
)

func SetRouter() {
	app := fiber.New()

	app.Get("/", controller.GetUser)

	app.Listen(":8080")
}
