package models

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"

	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/gomail.v2"
)

type Email struct {
	SMTPHost    string
	SMTPPort    int
	SenderEmail string
	SenderPass  string
}

type EmailTemplateData struct {
	Username  string
	ResetLink string
	AppName   string
	Year      int
}

func SendSignUpEmail(toEmail, username, token string) error {
	const smtpPort = 587
	smtpHost := config.AppConfig.MailSmtpHost
	smtpUser := config.AppConfig.MailSmtpUsername
	smtpPass := config.AppConfig.MailSmtpPassword
	appName := config.AppConfig.MailSender

	signUpURL := fmt.Sprintf("%s/confirm-signup?token=%s", config.AppConfig.FrontendHost, token)

	// Reading HTML template from file
	templatePath := filepath.Join("internal", "templates", "signUp.html")
	tmplBytes, err := os.ReadFile(templatePath)
	if err != nil {
		logger.Log.Printf("failed to read email template file: %v", err)
		return fmt.Errorf("failed to read email template file: %v", err)
	}

	tmpl, err := template.New("resetPassword").Parse(string(tmplBytes))
	if err != nil {
		logger.Log.Printf("failed to parse email template: %v", err)
		return fmt.Errorf("failed to parse email template: %v", err)
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, EmailTemplateData{
		Username:  username,
		ResetLink: signUpURL,
		AppName:   appName,
		Year:      time.Now().Year(),
	})
	if err != nil {
		logger.Log.Printf("failed to execute email template: %v", err)
		return fmt.Errorf("failed to execute email template: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser, appName+" Support")
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Confirm Account")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)

	if err := d.DialAndSend(m); err != nil {
		logger.Log.Printf("failed to send email: %v", err)
		return fmt.Errorf("failed to send email: %v", err)
	}

	logger.Log.Println("Sent email successfully!")
	return nil
}

func SendPasswordResetEmail(toEmail, username, token string) error {
	const smtpPort = 587
	smtpHost := config.AppConfig.MailSmtpHost
	smtpUser := config.AppConfig.MailSmtpUsername
	smtpPass := config.AppConfig.MailSmtpPassword
	appName := config.AppConfig.MailSender

	resetURL := fmt.Sprintf("%s/reset-password?token=%s", config.AppConfig.FrontendHost, token)

	// Reading HTML template from file
	templatePath := filepath.Join("internal", "templates", "passwordReset.html")
	tmplBytes, err := os.ReadFile(templatePath)
	if err != nil {
		logger.Log.Printf("failed to read email template file: %v", err)
		return fmt.Errorf("failed to read email template file: %v", err)
	}

	tmpl, err := template.New("resetPassword").Parse(string(tmplBytes))
	if err != nil {
		logger.Log.Printf("failed to parse email template: %v", err)
		return fmt.Errorf("failed to parse email template: %v", err)
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, EmailTemplateData{
		Username:  username,
		ResetLink: resetURL,
		AppName:   appName,
		Year:      time.Now().Year(),
	})
	if err != nil {
		logger.Log.Printf("failed to execute email template: %v", err)
		return fmt.Errorf("failed to execute email template: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser, appName+" Support")
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Reset your password")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)

	if err := d.DialAndSend(m); err != nil {
		logger.Log.Printf("failed to send email: %v", err)
		return fmt.Errorf("failed to send email: %v", err)
	}

	logger.Log.Println("Sent email successfully!")
	return nil
}
