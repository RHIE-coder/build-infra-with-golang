package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "Custom error message")
	})
	app.Listen(":8080")
}
