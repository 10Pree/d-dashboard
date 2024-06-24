package main

import (
	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("NONNY")
}

func main() {

	app := fiber.New()

	app.Get("/api", Welcome)

	app.Listen(":8080")
}
