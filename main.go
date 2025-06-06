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

			c.Status(502)

			return err

		}

		HasUser, err := db.CheckPresenceUser(user.Email)

		if err != nil {

			c.Status(502).SendString(err.Error())

			return err

		}

		switch HasUser {

		case true:

			msg = fmt.Sprintf("%s has already registered", user.Email)

			slog.Error(msg)

			c.Status(409).SendString(msg)

			return err

		default:

			email, err := db.CreateUser(user.FirstName, user.SecondName, user.Email, user.Password)

			if err != nil {

				c.Status(502).SendString(err.Error())

				return err

			}

			msg := fmt.Sprintf("Email %s was successfully registered", email)

			slog.Info(msg)

			c.Status(201).SendString(msg)

			return nil

		}

	})

	app.Post("/login", func(c *fiber.Ctx) error {

		user := new(structs.User)

		err := c.BodyParser(user)

		if err != nil {

			msg := "Error converting form data to struct"

			slog.Error(msg)

			return err

		}

		HasEmail, err := db.CheckPresenceUser(user.Email)

		if err != nil {

			c.Status(502).SendString(err.Error())

			return err

		}

		switch HasEmail {

		case true:

			correctPassword, err := db.CheckPassword(user.Email, user.Password)

			if err != nil {

				c.Status(502).SendString(err.Error())

				return err

			}

			if correctPassword {

				msg := fmt.Sprintf("Email %s was successfuly login", user.Email)

				slog.Info(msg)

				c.Status(200).SendString(msg)

				return nil

			}

			msg := "Password isn't corrected"

			slog.Error(msg)

			c.Status(401).SendString(msg)

			return err

		default:

			msg := fmt.Sprintf("Email %s isn't registered", user.Email)

			slog.Error(msg)

			c.Status(401).SendString(msg)

			return err

		}

	})

	app.Listen(":8001")

}
