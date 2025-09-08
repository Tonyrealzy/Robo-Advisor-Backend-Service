package services

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"

	"context"
	"fmt"

	sib "github.com/sendinblue/APIv3-go-library/v2/lib"
)

var apiClient *sib.APIClient

func InitEmailService() error {
	ctx := context.WithValue(context.Background(), sib.ContextAPIKey, sib.APIKey{
		Key: config.AppConfig.BrevoKey,
	})

	cfg := sib.NewConfiguration()
	cfg.BasePath = "https://api.brevo.com/v3"
	apiClient = sib.NewAPIClient(cfg)
	
	// var ctx context.Context
	// cfg := sib.NewConfiguration()
	// cfg.AddDefaultHeader("api-key", config.AppConfig.BrevoKey)
	// cfg.AddDefaultHeader("partner-key", config.AppConfig.BrevoKey)

	// apiClient = sib.NewAPIClient(cfg)
	result, resp, err := apiClient.AccountApi.GetAccount(ctx)
	if err != nil {
		logger.Log.Println("Error when calling AccountApi->get_account: ", err.Error())
		return fmt.Errorf("error when calling AccountApi->get_account: %v", err.Error())
	}

	logger.Log.Println("GetAccount Object:", result, " GetAccount Response: ", resp)
	return nil
}

func SendResetEmail(userEmail, userName, token string) error {
	sender := sib.SendSmtpEmailSender{Name: "Advisor Support", Email: config.AppConfig.MailSmtpUsername}
	to := []sib.SendSmtpEmailTo{{Email: userEmail, Name: userName}}

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", config.AppConfig.FrontendHost, token)

	result, httpResp, err := apiClient.TransactionalEmailsApi.
		SendTransacEmail(context.Background(), sib.SendSmtpEmail{
			Sender:     &sender,
			To:         to,
			TemplateId: 1,
			Params: map[string]interface{}{
				"reset_link": resetLink,
				"username":   userName,
			},
			Headers: map[string]interface{}{
				"X-Mailin-custom": "custom_header_1:custom_value_1|custom_header_2:custom_value_2",
			},
		})

	if err != nil {
		logger.Log.Println("Email send error:", err)
		return err
	}

	logger.Log.Printf("Email sent successfully from %v to %v", &sender, to)

	logger.Log.Println("Email send result:", result)
	logger.Log.Println("Email send http response:", httpResp)
	return nil
}
