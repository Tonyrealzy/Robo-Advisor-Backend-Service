package models

import (
	"go-backend/config"
	"time"
	"log"

	"gorm.io/gorm"
)

type UserSession struct {
	ID        string    `gorm:"type:uuid;primaryKey;unique;not null"`
	UserID    string    `gorm:"index;unique;not null"`
	Token     string    `gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
}

func (u *UserSession) CreateUserSession(db *gorm.DB, session *UserSession) error {
	err := config.CreateOneRecord(db, session)
	if err != nil {
		log.Printf("Error creating user session: %v", err)
		return err
	}

	return nil
}

func (u *UserSession) GetUserSessionByID(db *gorm.DB, userID string) (*UserSession, error) {
	var userSession UserSession

	err := config.FindOneByField(db, userSession, "user_id", userID)
	if err != nil {
		log.Printf("Error getting user session by ID: %v", err)
		return nil, err
	}

	return &userSession, nil
}

func (u *UserSession) GetUserSession(db *gorm.DB, userID string, token string) (*UserSession, error) {
	var userSession UserSession

	err := config.FindByTwoFields(db, userSession, "token", token, "user_id", userID)
	if err != nil {
		log.Printf("Error getting user session: %v", err)
		return nil, err
	}

	return &userSession, nil
}

func (u *UserSession) DeleteUserSession(db *gorm.DB, userID string) error {
	var userSession UserSession
	
	err := config.DeleteSpecificRecord(db, userSession, "user_id = ?", userID)
	if err != nil {
		log.Printf("Error deleting user session: %v", err)
		return err
	}

	return nil
}
