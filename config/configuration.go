package config

import (
	"os"
	"strings"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models/repository"

	"github.com/joho/godotenv"
)

var AppConfig repository.Config

func maskValue(val string) string {
	if len(val) <= 4 {
		return "****"
	}
	return val[:2] + strings.Repeat("*", len(val)-4) + val[len(val)-2:]
}

func LoadEnv() (repository.Config, error) {
	_, loadErr := os.Stat(".env")
	if loadErr == nil {
		err := godotenv.Load()
		if err != nil {
			logger.Log.Printf("Error loading .env file: %v", loadErr)
			return AppConfig, err
		}
	} else {
		logger.Log.Println(".env not found. Using platform environment variables.")
	}

	AppConfig = repository.Config{
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresSslMode:  os.Getenv("POSTGRES_SSLMODE"),
		JwtSecret:        os.Getenv("JWT_SECRET"),
		JwtExpiration:    os.Getenv("JWT_EXPIRATION"),
		AiService:        os.Getenv("AI_SERVICE"),
		ApiKey:           os.Getenv("GOOGLE_API_KEY"),
		FrontendHost:     os.Getenv("FRONTEND_HOST"),
		Port:             os.Getenv("PORT"),
		BrevoKey:         os.Getenv("BREVO_KEY"),
		AppEnv:           os.Getenv("APP_ENV"),
		MailSender:       os.Getenv("MAIL_SENDER"),
		MailSmtpHost:     os.Getenv("MAIL_SMTP_HOST"),
		MailSmtpUsername: os.Getenv("MAIL_SMTP_USERNAME"),
		MailSmtpPassword: os.Getenv("MAIL_SMTP_PASSWORD"),
	}

	required := map[string]string{
		"POSTGRES_HOST":      AppConfig.PostgresHost,
		"POSTGRES_PORT":      AppConfig.PostgresPort,
		"POSTGRES_USER":      AppConfig.PostgresUser,
		"POSTGRES_PASSWORD":  AppConfig.PostgresPassword,
		"POSTGRES_DB":        AppConfig.PostgresDB,
		"POSTGRES_SSLMODE":   AppConfig.PostgresSslMode,
		"JWT_SECRET":         AppConfig.JwtSecret,
		"JWT_EXPIRATION":     AppConfig.JwtExpiration,
		"AI_SERVICE":         AppConfig.AiService,
		"GOOGLE_API_KEY":     AppConfig.ApiKey,
		"FRONTEND_HOST":      AppConfig.FrontendHost,
		"PORT":               AppConfig.Port,
		"BREVO_KEY":          AppConfig.BrevoKey,
		"APP_ENV":            AppConfig.AppEnv,
		"MAIL_SENDER":        AppConfig.MailSender,
		"MAIL_SMTP_HOST":     AppConfig.MailSmtpHost,
		"MAIL_SMTP_USERNAME": AppConfig.MailSmtpUsername,
		"MAIL_SMTP_PASSWORD": AppConfig.MailSmtpPassword,
	}

	for key, val := range required {
		if val == "" {
			logger.Log.Fatalf("❌ Missing required environment variable: %s", key)
		} else {
			logger.Log.Printf("✅ %s is set (%s)", key, maskValue(val))
		}
	}

	logger.Log.Println("Loaded .env file successfully")

	return AppConfig, nil
}
