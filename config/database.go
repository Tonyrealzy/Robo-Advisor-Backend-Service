package config

import (
	"fmt"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB {
	connectedDB := connectToDb(AppConfig.PostgresHost,
		AppConfig.PostgresUser,
		AppConfig.PostgresPassword,
		AppConfig.PostgresDB,
		AppConfig.PostgresPort,
		AppConfig.PostgresSslMode,
	)

	logger.Log.Println("Database connected successfully")

	return connectedDB
}

func connectToDb(host, user, password, dbname, port, sslmode string) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Log.Fatal("Failed to connect to database:", err)
	}

	logger.Log.Printf("Connected to: %v", dbname)

	return db
}
