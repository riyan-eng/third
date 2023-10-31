package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Third World!")
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		fmt.Println(c.Params("name"))
		return c.SendString("Hello, Third World Name!")
	})

	app.Listen(":3000")
}
