// @title Robo-Advisor API
// @version 1.0
// @description This is a backend server for authentication and AI interaction.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"go-backend/config"
	"go-backend/middleware"
	"go-backend/models"
	"go-backend/routes"
	"log"
)

func main() {
	_, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Failed to load env credentials: %v", err)
	}

	db := config.ConnectToDatabase()
	if db != nil {
		log.Println("Ready to go!")
	}

	dbErr := db.AutoMigrate(&models.User{}, &models.PasswordReset{}, &models.UserSession{}, &models.AIPersistedResponse{})
	if dbErr != nil {
		log.Fatalf("Migration failed: %v", dbErr)
	} else {
		log.Println("Database auto-migrated successfully!")
	}

	router := middleware.SetupRouter()

	routes.SetupRoutes(router, db)

	port := config.AppConfig.Port
	if port == "" {
		port = "8080"
	}

	startErr := router.Run(":" + port)
	if startErr != nil {
		log.Fatalf("Server failed: %v", startErr)
	}
}
