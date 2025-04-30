package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectToDatabase() *gorm.DB {
	connectedDB := connectToDb(AppConfig.PostgresHost,
		AppConfig.PostgresUser,
		AppConfig.PostgresPassword,
		AppConfig.PostgresDB,
		AppConfig.PostgresPort,
		AppConfig.PostgresSslMode,
		AppConfig.PostgresTimezone,
	)
	
	log.Println("Database connected successfully")

	return connectedDB
}

func connectToDb(host, user, password, dbname, port, sslmode, timezone string) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", host, user, password, dbname, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Printf("Connected to: %v", dbname)

	return db
}
