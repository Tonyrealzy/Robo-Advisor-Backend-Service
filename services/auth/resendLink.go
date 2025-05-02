package auth

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/utils"
)

func SendLinkToUser(db *gorm.DB, existingUser *models.User) (string, error) {
	var password models.PasswordReset

	tokenString := fmt.Sprintf("%s-%s-%s", existingUser.Email, existingUser.ID, time.Now().UTC().String())
	hashedToken, err := utils.HashPassword(tokenString)
	if err != nil {
		logger.Log.Printf("error hashing token string: %v", err)
		return "", err
	}

	passwordReset := models.PasswordReset{
		ID:        utils.GenerateUUID(),
		UserID:    existingUser.ID,
		Email:     existingUser.Email,
		Token:     hashedToken,
		ExpiresAt: time.Now().UTC().Add(time.Minute * 30), // Token expires in 1/2 hour
	}
	createErr := password.CreatePasswordReset(db, &passwordReset)
	if createErr != nil {
		logger.Log.Printf("error resetting password: %v", createErr)
		return "", createErr
	}

	// Send a link with the reset token to the user's email
	emailErr := services.SendResetEmail(existingUser.Email, existingUser.Name, hashedToken)
	if emailErr != nil {
		return "", emailErr
	}
	
	return hashedToken, nil
}

func ResendLinkToUser(db *gorm.DB, email string) (string, error) {
	var password models.PasswordReset
	var user models.User

	reset, err := password.GetPasswordResetByEmail(db, email)
	if err != nil {
		logger.Log.Printf("error fetching password reset model: %v", err)
		return "", errors.New("error fetching password reset model")
	}
	existingUser, err := user.GetUserByEmail(db, email)
	if err != nil {
		logger.Log.Printf("error fetching password reset model: %v", err)
		return "", errors.New("error fetching password reset model")
	}

	_, expiryErr := password.GetPasswordResetByToken(db, reset.Token)
	if expiryErr != nil {
		tokenString := fmt.Sprintf("%s-%s-%s", reset.Email, reset.UserID, time.Now().UTC().String())
		hashedToken, err := utils.HashPassword(tokenString)
		if err != nil {
			logger.Log.Printf("error hashing token string: %v", err)
			return "", err
		}
		reset.Token = hashedToken
		reset.ExpiresAt = time.Now().UTC().Add(time.Minute * 30)
		updateErr := password.UpdatePasswordReset(db, reset)
		if updateErr != nil {
			logger.Log.Printf("error updating token string: %v", updateErr)
			return "", updateErr
		}
	}
	
	// Send a link with the reset token to the user's email
	emailErr := services.SendResetEmail(existingUser.Email, existingUser.Name, reset.Token)
	if emailErr != nil {
		return "", emailErr
	}

	return reset.Token, nil
}
