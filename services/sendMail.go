package services

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"

	"fmt"

	"gorm.io/gorm"
)

func SendEmail(db *gorm.DB, email []string, hashedToken string) error {
	emailConfig := models.Email{
		SMTPHost:    "smtp.gmail.com",
		SMTPPort:    587,
		SenderEmail: config.AppConfig.EmailAddress,  // your Gmail address
		SenderPass:  config.AppConfig.EmailPassword, // app password generated from Google
	}

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", config.AppConfig.FrontendHost, hashedToken)

	subject := "Password Reset Request"
	body := fmt.Sprintf(`
		<h2>Hello,</h2>
		<p>You requested a password reset. Please click the button below to reset your password:</p>
		<a href="%s" style="background-color: #4CAF50; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px;">Reset Password</a>
		<p>If you did not request this, you can ignore this email.</p>
		<p>This link will expire in 30 minutes.</p>
	`, resetLink)

	err := models.SendEmail(emailConfig, email, subject, body)
	if err != nil {
		return err
	}

	return nil
}
