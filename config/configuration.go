package config

import (
	"os"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models/repository"

	"github.com/joho/godotenv"
)

var AppConfig repository.Config

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
		FrontendHost:     os.Getenv("FRONTEND_HOST"),
		Port:             os.Getenv("PORT"),
		BrevoKey:         os.Getenv("BREVO_KEY"),
		AppEnv:           os.Getenv("APP_ENV"),
	}
	logger.Log.Println("Loaded .env file successfully")

	return AppConfig, nil
}
