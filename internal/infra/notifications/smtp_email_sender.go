package notifications

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type SMTPEmailSender struct {
	host     string
	port     string
	email    string
	password string
}

func NewSMTPEmailSender() *SMTPEmailSender {
	godotenv.Load()

	return &SMTPEmailSender{
		host:     os.Getenv("EMAIL_HOST"),
		port:     os.Getenv("EMAIL_PORT"),
		email:    os.Getenv("EMAIL_ADDRESS"),
		password: os.Getenv("EMAIL_PASSWORD"),
	}
}

func (s *SMTPEmailSender) Send(to string, subject string, body string) error {
	godotenv.Load()
	os.Getenv("JWT_SECRET")

	auth := smtp.PlainAuth("", s.email, s.password, s.host)

	msg := []byte(fmt.Sprintf("Subject: %v\r\n\r\n%v", subject, body))

	return smtp.SendMail(s.host+":"+s.port, auth, s.email, []string{to}, msg)
}
