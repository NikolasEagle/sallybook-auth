package main

import (
	db "sallybook-auth/funcs"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("SallyBook")

	})

	app.Get("/check", func(c *fiber.Ctx) error {

		db.CheckConnection()

		return nil

	})

	app.Listen(":8001")

}
