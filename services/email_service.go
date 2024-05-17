package services

import (
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

func SendEmail(to string, subject string, body string) error {
    e := email.NewEmail()
    e.From = os.Getenv("EMAIL_FROM")
    e.To = []string{to}
    e.Subject = subject
    e.Text = []byte(body)
    return e.Send("smtp.gmail.com:587", smtp.PlainAuth("", os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"), "smtp.gmail.com"))
}
