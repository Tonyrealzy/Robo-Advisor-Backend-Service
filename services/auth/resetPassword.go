package auth

import (
	"fmt"
	"time"
	"gorm.io/gorm"

	"github.com/Tonyrealzy/go-backend/models"
	"github.com/Tonyrealzy/go-backend/services"
	"github.com/Tonyrealzy/go-backend/utils"
)

func ResetPassword(db *gorm.DB, email string) (string, error) {
	var user *models.User
	var password *models.PasswordReset

	existingUser, err := user.GetUserByEmail(db, email)
	if err != nil {
		return "", fmt.Errorf("user not found: %v", err)
	}

	tokenString := fmt.Sprintf("%s-%s-%s", email, existingUser.ID, time.Now().String())
	hashedToken, err := utils.HashPassword(tokenString)
	if err != nil {
		return "", err
	}

	passwordReset := models.PasswordReset{
		ID:        utils.GenerateUUID(),
		UserID:    existingUser.ID,
		Token:     hashedToken,
		ExpiresAt: time.Now().Add(time.Minute * 30), // Token expires in 1/2 hour
	}
	createErr := password.CreatePasswordReset(db, &passwordReset)
	if createErr != nil {
		return "", createErr
	}

	// You would send a link with the reset token to the user's email
	emailErr := services.SendEmail(db, []string{email}, hashedToken)
	if emailErr != nil {
		return "", emailErr
	}

	return hashedToken, nil
}
