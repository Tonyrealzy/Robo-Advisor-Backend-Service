package auth

import (
	"errors"
	"fmt"
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
		return "", fmt.Errorf("user not found: %v", err)
	}

	userExists := utils.CheckPasswordHash(password, existingUser.Password)
	if !userExists {
		return "", errors.New("incorrect password")
	}

	token, err := middleware.CreateToken(existingUser.ID, existingUser.Email)
	if err != nil {
		return "", err
	}

	userToken := models.UserSession{
		UserID:    existingUser.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Minute * 30),
	}

	createErr := session.CreateUserSession(db, &userToken)
	if createErr != nil {
		return "", createErr
	}

	return token, nil
}
