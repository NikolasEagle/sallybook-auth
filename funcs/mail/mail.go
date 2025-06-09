package mail

import (
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

var (
	smtpHost = os.Getenv("SMTP_HOST")

	smtpPort = os.Getenv("SMTP_PORT")

	smtpUser = os.Getenv("SMTP_USER")

	smtpPassword = os.Getenv("SMTP_PASSWORD")

	smtpRecipient = os.Getenv("SMTP_RECIPIENT")
)
