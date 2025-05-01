package auth

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
)

func ConfirmSignup(db *gorm.DB, email, token string) (string, error) {
	var user models.User
	var reset models.PasswordReset

	existingUser, err := user.GetUserByEmail(db, email)
	if err != nil {
		return "", fmt.Errorf("user not found: %v", err)
	}

	searchErr := config.FindByTwoFields(db, &reset, "token = ?", token, "expires_at > ?", time.Now())
	if searchErr != nil {
		return "", fmt.Errorf("invalid or expired reset token")
	}

	userErr := config.FindByID(db, &user, reset.UserID)
	if userErr != nil {
		return "", fmt.Errorf("user not found")
	}

	existingUser.IsActive = true
	updateErr := config.UpdateModel(db, &existingUser)
	if updateErr != nil {
		return "", fmt.Errorf("failed to update user status: %v", updateErr)
	}

	return "User status updated successfully!", nil
}
