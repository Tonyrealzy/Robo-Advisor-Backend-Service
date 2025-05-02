package auth

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
)

func ConfirmSignup(db *gorm.DB, email, token string) (string, error) {
	var user models.User
	var reset models.PasswordReset

	existingUser, err := user.GetUserByEmail(db, email)
	if err != nil {
		logger.Log.Printf("user not found: %v", err)
		return "", fmt.Errorf("user not found: %v", err)
	}

	searchRes, searchErr := reset.GetPasswordResetByToken(db, token)
	if searchErr != nil {
		logger.Log.Printf("invalid or expired reset token: %v", searchErr)
		return "", fmt.Errorf("invalid or expired reset token")
	}

	userById, userErr := user.GetUserByID(db, searchRes.UserID)
	if userErr != nil {
		logger.Log.Printf("user not found: %v", userErr)
		return "", fmt.Errorf("user not found")
	}

	if existingUser.ID == userById.ID {
		existingUser.IsActive = true
		updateErr := config.UpdateModel(db, &existingUser)
		if updateErr != nil {
			logger.Log.Printf("failed to update user status: %v", updateErr)
			return "", fmt.Errorf("failed to update user status: %v", updateErr)
		}
	} else {
		return "", fmt.Errorf("invalid reset token")
	}

	return "User status updated successfully!", nil
}
