package models

import (
	"fmt"
	"time"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"

	"gorm.io/gorm"
)

type PasswordReset struct {
	ID        string    `gorm:"type:uuid;primaryKey;unique;not null"`
	UserID    string    `gorm:"index"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Token     string    `gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
}

type PasswordResetRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type PasswordChangeRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

type ConfirmSignupRequest struct {
	Token string `json:"token" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (p *PasswordReset) CreatePasswordReset(db *gorm.DB, passReset *PasswordReset) error {
	err := config.CreateOneRecord(db, passReset)
	if err != nil {
		logger.Log.Printf("Error creating password reset model: %v", err)
		return err
	}

	return nil
}

func (p *PasswordReset) GetPasswordResetByToken(db *gorm.DB, token string) (*PasswordReset, error) {
	var passReset PasswordReset

	result := db.Where(fmt.Sprintf("%s = ? AND %s > ?", "token", "expires_at"), token, time.Now().UTC()).First(&passReset)
	if result.Error != nil {
		logger.Log.Printf("invalid or expired reset token: %v", result.Error)
		return nil, result.Error
	}

	return &passReset, nil
}

func (p *PasswordReset) GetPasswordResetByEmail(db *gorm.DB, email string) (*PasswordReset, error) {
	var passReset PasswordReset

	err := config.FindOneByField(db, &passReset, "email", email)
	if err != nil {
		logger.Log.Printf("invalid or expired reset token: %v", err)
		return nil, err
	}

	return &passReset, nil
}

func (p *PasswordReset) GetPasswordResetByID(db *gorm.DB, id string) (*PasswordReset, error) {
	var reset PasswordReset

	err := config.FindByID(db, &reset, id)
	if err != nil {
		logger.Log.Printf("Error getting reset model by ID: %v", err)
		return nil, err
	}

	return &reset, nil
}

func (p *PasswordReset) UpdatePasswordReset(db *gorm.DB, reset *PasswordReset) error {
	err := config.UpdateModel(db, &reset)
	if err != nil {
		logger.Log.Printf("Error updating token in password reset model: %v", err)
		return err
	}

	return nil
}
