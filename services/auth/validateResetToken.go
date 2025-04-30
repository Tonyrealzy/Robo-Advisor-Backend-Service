package auth

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/utils"

	"fmt"
	"time"

	"gorm.io/gorm"
)

func ValidateResetToken(db *gorm.DB, token, password string) (string, error) {
	var reset models.PasswordReset
	var user models.User

	err := config.FindByTwoFields(db, &reset, "token = ?", token, "expires_at > ?", time.Now())
	if err != nil {
		return "", fmt.Errorf("invalid or expired reset token")
	}

	userErr := config.FindByID(db, &user, reset.UserID)
	if userErr != nil {
		return "", fmt.Errorf("user not found")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword
	updateErr := config.UpdateModel(db, &user)
	if updateErr != nil {
		return "", fmt.Errorf("failed to update password: %v", updateErr)
	}

	deleteErr := config.DeleteByID(db, &reset, reset.ID)
	if deleteErr != nil {
		return "", fmt.Errorf("failed to delete reset token: %v", deleteErr)
	}

	return fmt.Sprint("Password reset successful"), nil
}
