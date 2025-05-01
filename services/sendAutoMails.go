package services

import (
	"fmt"
	"os"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	mailjet "github.com/mailjet/mailjet-apiv3-go/v4"
)

func SendResetEmail(toEmail, toName, resetLink string) error {
	mj := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

	messageInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "your-sender@example.com",
				Name:  "Your App Name",
			},
			To: &mailjet.RecipientsV31{
				{
					Email: toEmail,
					Name:  toName,
				},
			},
			Subject:  "Password Reset Request",
			TextPart: "Reset your password using the link below.",
			HTMLPart: fmt.Sprintf(`<p>Hello %s,</p><p>Click <a href="%s">here</a> to reset your password. This link will expire in 15 minutes.</p>`, toName, resetLink),
			CustomID: "PasswordResetEmail",
		},
	}

	messages := mailjet.MessagesV31{Info: messageInfo}

	res, err := mj.SendMailV31(&messages)
	if err != nil {
		logger.Log.Printf("error sending email: %v", err)
		return fmt.Errorf("error sending email: %w", err)
	}

	logger.Log.Printf("Email sent! Response: %+v\n", res)
	return nil
}
