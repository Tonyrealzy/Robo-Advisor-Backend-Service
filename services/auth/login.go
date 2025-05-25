package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/middleware"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/utils"

	"gorm.io/gorm"
)

func Login(db *gorm.DB, email, password string) (string, error) {
	var user models.User
	var session models.UserSession

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

	userIsActive, userActiveErr := user.IsUserActive(db, existingUser.IsActive)
	if userActiveErr != nil {
		logger.Log.Printf("Error checking if user is active: %v", userActiveErr)
		return "", userActiveErr
	}
	if !userIsActive {
		logger.Log.Printf("User is inactive. Resending sign-up confirmation link.")
		_, messageErr := ResendLinkToUser(db, email)
		if messageErr != nil {
			return "", messageErr
		}
		return "", errors.New("user is inactive. check mail for activation link")
	}

	token, err := middleware.CreateToken(existingUser.ID, existingUser.Email)
	if err != nil {
		logger.Log.Printf("Error creating token: %v", err)
		return "", err
	}

	sessionFetched, sessionErr := session.GetUserSessionByID(db, existingUser.ID)
	if sessionErr != nil && !errors.Is(sessionErr, gorm.ErrRecordNotFound) {
		logger.Log.Printf("Error retrieving user session: %v", sessionErr)
		return "", sessionErr
	}

	if sessionFetched == nil || sessionFetched.Token == "" {
		userToken := models.UserSession{
			ID:        utils.GenerateUUID(),
			UserID:    existingUser.ID,
			Token:     token,
			ExpiresAt: time.Now().UTC().Add(time.Minute * 30),
		}

		createErr := session.CreateUserSession(db, &userToken)
		if createErr != nil {
			logger.Log.Printf("Error creating user session: %v", createErr)
			return "", createErr
		}
	} else {
		sessionFetched.Token = token
		sessionFetched.ExpiresAt = time.Now().UTC().Add(time.Minute * 30)

		err := session.UpdateUserSession(db, sessionFetched)
		if err != nil {
			logger.Log.Printf("Error updating user session: %v", err)
			return "", err
		}
	}

	return token, nil
}
