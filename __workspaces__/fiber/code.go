package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func SIMPLE_SERVER_RUNNER() {
	app := fiber.New()

	log.Fatal(app.Listen(":5000"))
}

func main() {

}
