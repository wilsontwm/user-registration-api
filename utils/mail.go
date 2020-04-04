package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

func Email(name string, email string, subject string, plainContent string, htmlContent string) error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Error loading .env file.")
	}

	from := mail.NewEmail(os.Getenv("app_name"), "hello@example.com")
	to := mail.NewEmail(name, email)
	message := mail.NewSingleEmail(from, subject, to, plainContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("sendgrid_api_key"))
	_, err = client.Send(message)

	return err
}
