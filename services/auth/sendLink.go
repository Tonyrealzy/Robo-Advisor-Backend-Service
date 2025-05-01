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

func SendLinkToUser(db *gorm.DB, existingUser *models.User, email string) (string, error) {
	var password *models.PasswordReset

	tokenString := fmt.Sprintf("%s-%s-%s", email, existingUser.ID, time.Now().String())
	hashedToken, err := utils.HashPassword(tokenString)
	if err != nil {
		logger.Log.Printf("error hashing password: %v", err)
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
		logger.Log.Printf("error resetting password: %v", createErr)
		return "", createErr
	}

	// You would send a link with the reset token to the user's email
	// emailErr := services.SendEmail(db, []string{email}, hashedToken)
	// if emailErr != nil {
	// 	return "", emailErr
	// }

	return hashedToken, nil
}
