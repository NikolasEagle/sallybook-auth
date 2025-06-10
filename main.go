package main

import (
	"fmt"
	"log/slog"
	"os"
	"sallybook-auth/funcs/db"
	"sallybook-auth/funcs/mail"
	"sallybook-auth/funcs/redis_store"
	"sallybook-auth/structs"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

func main() {

	app := fiber.New()

	storage := redis_store.Store

	store := session.New(session.Config{

		Storage: storage,

		Expiration: 7 * 24 * time.Hour,

		KeyLookup: "cookie:session_id",

		KeyGenerator: utils.UUIDv4,
	})

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

		user := new(structs.User)

		err := c.BodyParser(user)

		if err != nil {

			slog.Error("Error converting form data to struct")

			c.Status(502).SendString("Error data processing")

			return err

		}

		hasUser, err := db.CheckPresenceUser(user.Email)

		if err != nil {

			c.Status(502).SendString(err.Error())

			return err

		}

		switch hasUser {

		case true:

			msg := fmt.Sprintf("%s has already registered", user.Email)

			slog.Error(msg)

			c.Status(409).SendString(msg)

			return err

		default:

			email, err := db.CreateUser(user.FirstName, user.SecondName, user.Email, user.Password)

			if err != nil {

				c.Status(502).SendString(err.Error())

				return err

			}

			err = mail.SendMessageToAdmin(user.FirstName, user.SecondName, user.Email, user.Password)

			if err != nil {

				c.Status(502).SendString(err.Error())

				return err

			}

			err = mail.SendMessageToUser(user.FirstName, user.SecondName, user.Email, user.Password)

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

		sess, err := store.Get(c)

		if err != nil {

			slog.Error("Error initialization session storage")

			c.Status(502).SendString("Error data processing")

			return err

		}

		user := new(structs.User)

		err = c.BodyParser(user)

		if err != nil {

			slog.Error("Error converting form data to struct")

			c.Status(502).SendString("Error data processing")

			return err

		}

		correctEmail, err := db.CheckPresenceUser(user.Email)

		if err != nil {

			c.Status(502).SendString(err.Error())

			return err

		}

		switch correctEmail {

		case true:

			correctPassword, err := db.CheckPassword(user.Email, user.Password)

			if err != nil {

				c.Status(502).SendString(err.Error())

				return err

			}

			if correctPassword {

				sess.Set("email", user.Email)

				err := sess.Save()

				if err != nil {

					slog.Error(fmt.Sprintf("Error saving session for %s", user.Email))

					c.Status(502).SendString("Error creating session")

					return err

				}

				msg := fmt.Sprintf("Email %s was successfuly login", user.Email)

				slog.Info(msg)

				c.Status(200).SendString(msg)

				return nil

			}

			msg := "Password isn't corrected"

			slog.Error(msg)

			c.Status(401).SendString(msg)

			return nil

		default:

			msg := fmt.Sprintf("Email %s isn't registered", user.Email)

			slog.Error(msg)

			c.Status(401).SendString(msg)

			return nil

		}

	})

	app.Post("/check_auth", func(c *fiber.Ctx) error {

		sess, err := store.Get(c)

		if err != nil {

			slog.Error("Error initialization session storage")

			c.Status(502).SendString("Error data processing")

			return err

		}

		hasSession, err := redis_store.GetValue(sess.ID())

		if err != nil {

			c.Status(502).SendString("Error data processing")

			return err

		}

		switch hasSession {

		case "none":

			msg := fmt.Sprintf("Email %s wasn't login", sess.Get("email"))

			slog.Error(msg)

			c.Status(401).SendString(msg)

			return nil

		default:

			msg := fmt.Sprintf("Email %s was successfuly login", sess.Get("email"))

			slog.Info(msg)

			c.Status(200).SendString(msg)

			return nil

		}

	})

	app.Post("/logout", func(c *fiber.Ctx) error {

		sess, err := store.Get(c)

		if err != nil {

			slog.Error("Error initialization session storage")

			c.Status(502).SendString("Error data processing")

			return err

		}

		err = sess.Destroy()

		if err != nil {

			slog.Error("Error destroying session storage")

			c.Status(502).SendString("Error data processing")

			return err

		}

		return nil

	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT_SERVER_AUTH")))

}
