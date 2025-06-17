// @title Robo-Advisor API
// @version 1.0
// @description This is a backend server for authentication and AI interaction.
// @host robo-advisor-backend-service.onrender.com
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/middleware"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/routes"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Starting Robo Advisor Backend Service")

	repo, err := config.LoadEnv()
	if err != nil {
		logger.Log.Errorf("Failed to load env credentials: %v", err)
	}

	client, err := models.NewAIService(repo.ApiKey)
	if err != nil {
		logger.Log.Errorf("Failed to load AI Service: %v", err)
	}

	if err := services.InitEmailService(); err != nil {
		logger.Log.Fatalf("Failed to initialise email service: %v", err)
	}

	db := config.ConnectToDatabase()
	if db != nil {
		logger.Log.Println("Ready to go!")
	}

	if config.AppConfig.AppEnv == "development" {
		err := db.AutoMigrate(&models.User{}, &models.PasswordReset{}, &models.UserSession{}, &models.AIPersistedResponse{})
		if err != nil {
			logger.Log.Fatalf("Migration failed: %v", err)
		} else {
			logger.Log.Println("Database auto-migrated successfully!")
		}
	}

	router := middleware.SetupRouter()

	routes.SetupRoutes(router, db, client)

	port := config.AppConfig.Port
	if port == "" {
		port = "8080"
	}

	startErr := router.Run(":" + port)
	if startErr != nil {
		logger.Log.Fatalf("Server failed: %v", startErr)
	}
}
