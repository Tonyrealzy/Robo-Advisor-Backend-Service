package auth

import (
	"errors"
	"fmt"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/middleware"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/utils"
	"time"

	"gorm.io/gorm"
)

func Login(db *gorm.DB, email, password string) (string, error) {
	var user *models.User
	var session *models.UserSession

	existingUser, err := user.GetUserByEmail(db, email)
	if err != nil {
		logger.Log.Printf("user not found: %v", err)
		return "", fmt.Errorf("user not found: %v", err)
	}

	userExists := utils.CheckPasswordHash(password, existingUser.Password)
	if !userExists {
		logger.Log.Printf("incorrect password")
		return "", errors.New("incorrect password")
	}

	token, err := middleware.CreateToken(existingUser.ID, existingUser.Email)
	if err != nil {
		logger.Log.Printf("Error creating token: %v", err)
		return "", err
	}

	userToken := models.UserSession{
		UserID:    existingUser.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Minute * 30),
	}

	createErr := session.CreateUserSession(db, &userToken)
	if createErr != nil {
		logger.Log.Printf("Error creating user session: %v", createErr)
		return "", createErr
	}

	return token, nil
}
