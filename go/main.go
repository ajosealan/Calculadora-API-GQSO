package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("CALCULADORA")
	})

	app.Get("/soma/:op1/:op2", soma)

	app.Listen(":8000")
}
