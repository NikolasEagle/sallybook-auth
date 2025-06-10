package mail

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

var _ = godotenv.Load()

var (
	smtpHost = os.Getenv("SMTP_HOST")

	smtpPort = os.Getenv("SMTP_PORT")

	smtpUser = os.Getenv("SMTP_USER")

	smtpPassword = os.Getenv("SMTP_PASSWORD")

	smtpRecipient = os.Getenv("SMTP_RECIPIENT")
)

func SendMessageToAdmin(first_name, second_name, email, password string) {

	message := gomail.NewMessage()

	message.SetHeader("From", fmt.Sprintf("MouseBook <%s>", smtpUser))

	message.SetHeader("To", smtpRecipient)

	message.SetHeader("Subject", "Зарегистрирован новый пользователь")

	message.SetBody("text/html", fmt.Sprintf(`
	
		<h1>NEW USER</h1>
        <p>Name: %s</p>
        <p>Surname: %s</p>
        <p><b>Email:%s</p>
        <p><b>Password:</b> %s</p>
	
	`, first_name, second_name, email, password))

}
