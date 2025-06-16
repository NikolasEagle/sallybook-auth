package mail

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/metaer/go-easy-dkim-signer/easydkim"
	"gopkg.in/gomail.v2"
)

var _ = godotenv.Load()

var (
	smtpHost = os.Getenv("SMTP_HOST")

	smtpPort = os.Getenv("SMTP_PORT")

	smtpUser = os.Getenv("SMTP_USER")

	/*smtpPassword = os.Getenv("SMTP_PASSWORD")*/

	smtpRecipient = os.Getenv("SMTP_RECIPIENT")

	domain = os.Getenv("DOMAIN")

	dkimSelector = os.Getenv("DKIM_SELECTOR")

	dkimPrivateFileKey = os.Getenv("DKIM_PRIVATE_FILE_KEY")
)

func SendMessageToAdmin(first_name, second_name, email, password string) error {

	message := gomail.NewMessage()

	message.SetHeader("From", fmt.Sprintf("MouseBook <%s>", smtpUser))

	message.SetHeader("To", smtpRecipient)

	message.SetHeader("Subject", "Зарегистрирован новый пользователь")

	message.SetBody("text/html", fmt.Sprintf(`
	
		<h1>NEW USER</h1>
        <p>Name: %s</p>
        <p>Surname: %s</p>
        <p><b>Email:</b> %s</p>
        <p><b>Password:</b> %s</p>
	
	`, first_name, second_name, email, password))

	var buffer bytes.Buffer

	_, err := message.WriteTo(&buffer)

	if err != nil {

		slog.Error(err.Error())

		return fmt.Errorf("%s", "Error converting data mail message")

	}

	var signedMessage []byte

	signedMessage, err = easydkim.Sign(buffer.Bytes(), dkimPrivateFileKey, dkimSelector, domain)

	if err != nil {

		slog.Error(err.Error())

		return fmt.Errorf("%s", "Error signing data mail message")

	}

	/*auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)*/

	err = smtp.SendMail(fmt.Sprintf("%s:%s", smtpHost, smtpPort), nil, smtpUser, []string{smtpRecipient}, signedMessage)

	if err != nil {

		slog.Error(err.Error())

		return fmt.Errorf("%s", "Error sending mail message")

	}

	return nil

}

func SendMessageToUser(first_name, second_name, email, password string) error {

	message := gomail.NewMessage()

	message.SetHeader("From", fmt.Sprintf("MouseBook <%s>", smtpUser))

	message.SetHeader("To", email)

	message.SetHeader("Subject", "Поздравляю с регистрацией в сервисе MouseBook")

	message.SetBody("text/html", fmt.Sprintf(`
	
		<h1>Данные для входа</h1>

        <p><b>Email:</b> %s</p>
        <p><b>Пароль:</b> %s</p>
	
	`, email, password))

	var buffer bytes.Buffer

	_, err := message.WriteTo(&buffer)

	if err != nil {

		slog.Error(err.Error())

		return fmt.Errorf("%s", "Error converting data mail message")

	}

	var signedMessage []byte

	signedMessage, err = easydkim.Sign(buffer.Bytes(), dkimPrivateFileKey, dkimSelector, domain)

	if err != nil {

		slog.Error(err.Error())

		return fmt.Errorf("%s", "Error signing data mail message")

	}

	/*auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)*/

	err = smtp.SendMail(fmt.Sprintf("%s:%s", smtpHost, smtpPort), nil, smtpUser, []string{email}, signedMessage)

	if err != nil {

		slog.Error(err.Error())

		return fmt.Errorf("%s", "Error sending mail message")

	}

	return nil

}
