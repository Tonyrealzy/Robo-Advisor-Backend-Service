package models

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"gopkg.in/gomail.v2"
)

type Email struct {
	SMTPHost    string
	SMTPPort    int
	SenderEmail string
	SenderPass  string
}

func SendEmail(cfg Email, to []string, subject, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", cfg.SenderEmail)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SenderEmail, cfg.SenderPass)

	if err := d.DialAndSend(m); err != nil {
		logger.Log.Printf("Error sending mail: %v", err)
		return err
	}

	return nil
}
