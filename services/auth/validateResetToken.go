package auth

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/utils"

	"fmt"
	"time"

	"gorm.io/gorm"
)

func ValidateResetToken(db *gorm.DB, token, password string) (string, error) {
	var reset models.PasswordReset
	var user models.User

	err := config.FindByTwoFields(db, &reset, "token = ?", token, "expires_at > ?", time.Now().UTC())
	if err != nil {
		logger.Log.Printf("error: %v", err)
		return "", fmt.Errorf("invalid or expired reset token")
	}

	userErr := config.FindByID(db, &user, reset.UserID)
	if userErr != nil {
		logger.Log.Printf("user not found: %v", userErr)
		return "", fmt.Errorf("user not found")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		logger.Log.Printf("error hashing password: %v", err)
		return "", err
	}

	user.Password = hashedPassword
	updateErr := config.UpdateModel(db, &user)
	if updateErr != nil {
		logger.Log.Printf("failed to update password: %v", updateErr)
		return "", fmt.Errorf("failed to update password")
	}

	deleteErr := config.DeleteSpecificRecord(db, reset, "id = ?", reset.ID)
	if deleteErr != nil {
		logger.Log.Printf("failed to delete reset token: %v", deleteErr)
		return "", fmt.Errorf("failed to delete reset token")
	}

	return "Password reset successful", nil
}
