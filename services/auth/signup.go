package auth

import (
	"errors"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/utils"

	"gorm.io/gorm"
)

func Signup(db *gorm.DB, email, password, firstName, lastName, userName string) (*models.User, string, error) {
	var user models.User

	existingUser, err := user.GetUserByEmail(db, email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Log.Errorf("database error checking user: %v", err)
		return nil, "", errors.New("database error checking user")
	}
	if err == nil && existingUser.Email != "" {
		logger.Log.Warn("email already in use")
		return nil, "", errors.New("email already in use")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		logger.Log.Printf("error hashing password: %v", err)
		return nil, "", err
	}

	newUser := models.User{
		ID:        utils.GenerateUUID(),
		Name:      userName,
		Email:     email,
		Password:  hashedPassword,
		FirstName: firstName,
		LastName:  lastName,
	}

	createErr := user.CreateUser(db, &newUser)
	if createErr != nil {
		logger.Log.Printf("error creating user: %v", createErr)
		return nil, "", createErr
	}

	linkMsg, linkErr := SendSignUpLinkToUser(db, &newUser)
	if linkErr != nil {
		return nil, "", linkErr
	}

	return &newUser, linkMsg, nil
}
