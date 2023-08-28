package main

import (
	"fmt"
	"reflect"
)

// import "github.com/gofiber/fiber/v2"

func AAA(a interface{}) {
	b := reflect.TypeOf(a)
	fmt.Println(b.Name() == "int32")
}

func main() {
	// app := fiber.New()
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return fiber.NewError(782, "Custom error message")
	// })
	// app.Listen(":8080")
	var a int32
	AAA(a)
}
