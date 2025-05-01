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
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Starting Robo Advisor Backend Service")

	_, err := config.LoadEnv()
	if err != nil {
		logger.Log.Errorf("Failed to load env credentials: %v", err)
	}

	db := config.ConnectToDatabase()
	if db != nil {
		logger.Log.Println("Ready to go!")
	}

	dbErr := db.AutoMigrate(&models.User{}, &models.PasswordReset{}, &models.UserSession{}, &models.AIPersistedResponse{})
	if dbErr != nil {
		logger.Log.Fatalf("Migration failed: %v", dbErr)
	} else {
		logger.Log.Println("Database auto-migrated successfully!")
	}

	router := middleware.SetupRouter()

	routes.SetupRoutes(router, db)

	port := config.AppConfig.Port
	if port == "" {
		port = "8080"
	}

	startErr := router.Run(":" + port)
	if startErr != nil {
		logger.Log.Fatalf("Server failed: %v", startErr)
	}
}
