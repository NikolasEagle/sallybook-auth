package main

import (
	"fmt"
	"log/slog"
	"sallybook-auth/funcs/db"
	"sallybook-auth/structs"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("SallyBook")

	})

	app.Get("/check", func(c *fiber.Ctx) error {

		_, err := db.CheckConnection()

		if err != nil {

			c.Status(502).SendString(err.Error())

			return err

		}

		msg := "Successfully connected to database!"

		slog.Info(msg)

		c.Status(200).SendString(msg)

		return nil

	})

	app.Post("/register", func(c *fiber.Ctx) error {

		var msg string

		user := new(structs.User)

		err := c.BodyParser(user)

		if err != nil {

			msg = "Error converting form data to struct"

			slog.Error(msg)

			return err

		}

		HasUser, err := db.CheckPresenceUser(user.Email)

		if err != nil {

			slog.Error(err.Error())

			return err

		}

		switch HasUser {

		case true:

			msg = fmt.Sprintf("%s has already registered", user.Email)

			c.Status(409).SendString(msg)

			slog.Error(msg)

			return err

		default:

			return nil

		}

	})

	app.Post("/login", func(c *fiber.Ctx) error {

		return nil

	})

	app.Listen(":8001")

}
