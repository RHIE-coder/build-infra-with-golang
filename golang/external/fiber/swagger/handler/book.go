package handler

import "github.com/gofiber/fiber/v2"

func AddBook(ctx *fiber.Ctx) error {
	return ctx.SendString("AddBook")
}

func ShowBook(ctx *fiber.Ctx) error {
	return ctx.SendString("ShowBook")
}

func ListBook(ctx *fiber.Ctx) error {
	return ctx.SendString("ListBook")
}

func UpdateBook(ctx *fiber.Ctx) error {
	return ctx.SendString("UpdateBook")
}

func DeleteBook(ctx *fiber.Ctx) error {
	return ctx.SendString("DeleteBook")
}
