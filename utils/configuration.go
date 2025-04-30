package utils

import (
	"log"
	"os"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models/repository"

	"github.com/joho/godotenv"
)

var AppConfig repository.Config

func LoadEnv() (repository.Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
		return AppConfig, err
	}

	AppConfig = repository.Config{
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresTimezone: os.Getenv("POSTGRES_TIMEZONE"),
		PostgresSslMode:  os.Getenv("POSTGRES_SSLMODE"),
		JwtSecret:        os.Getenv("JWT_SECRET"),
		JwtExpiration:    os.Getenv("JWT_EXPIRATION"),
		AiService:        os.Getenv("AI_SERVICE"),
		FrontendHost:     os.Getenv("FRONTEND_HOST"),
		Port:             os.Getenv("PORT"),
		EmailAddress:     os.Getenv("EMAIL_ADDRESS"),
		EmailPassword:    os.Getenv("EMAIL_PASSWORD"),
	}

	return AppConfig, nil
}
