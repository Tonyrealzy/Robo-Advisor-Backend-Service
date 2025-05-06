package auth

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	// "github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/utils"
)

func ResetPassword(db *gorm.DB, email string) (string, error) {
	var user models.User
	var password models.PasswordReset

	existingUser, err := user.GetUserByEmail(db, email)
	if err != nil {
		logger.Log.Printf("user not found: %v", err)
		return "", fmt.Errorf("user not found: %v", err)
	}

	tokenString := fmt.Sprintf("%s-%s-%s", email, existingUser.ID, time.Now().UTC().String())
	hashedToken, err := utils.HashPassword(tokenString)
	if err != nil {
		logger.Log.Printf("Error hashing password: %v", err)
		return "", err
	}

	resetGotten, resetErr := password.GetPasswordResetByEmail(db, email)
	if resetErr != nil {
		logger.Log.Printf("Error fetching password reset model by email: %v", resetErr)
		return "", resetErr
	}

	if resetGotten.Email == "" {
		passwordReset := models.PasswordReset{
			ID:        utils.GenerateUUID(),
			Email:     email,
			UserID:    existingUser.ID,
			Token:     hashedToken,
			ExpiresAt: time.Now().UTC().Add(time.Minute * 30), // Token expires in 1/2 hour
		}
		createErr := password.CreatePasswordReset(db, &passwordReset)
		if createErr != nil {
			logger.Log.Printf("Error resetting password: %v", createErr)
			return "", createErr
		}

	} else {
		resetGotten.Token = hashedToken
		resetGotten.ExpiresAt = time.Now().UTC().Add(time.Minute * 30)
		updateErr := password.UpdatePasswordReset(db, resetGotten)
		if updateErr != nil {
			logger.Log.Printf("error updating token string: %v", updateErr)
			return "", updateErr
		}
	}

	// You would send a link with the reset token to the user's email
	emailErr := models.SendPasswordResetEmail(existingUser.Email, existingUser.Name, hashedToken)
	if emailErr != nil {
		return "", emailErr
	}

	return hashedToken, nil
}
