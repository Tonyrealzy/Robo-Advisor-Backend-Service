package auth

import (
	"errors"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/utils"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"

	"gorm.io/gorm"
)

func Signup(db *gorm.DB, email, password, firstName, lastName, userName string) (*models.User, error) {
	var user *models.User

	existingUser, _ := user.GetUserByEmail(db, email)
	if existingUser.Name != "" {
		logger.Log.Printf("email already in use")
		return nil, errors.New("email already in use")
	}
	
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		logger.Log.Printf("error hashing password: %v", err)
		return nil, err
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
		return nil, createErr
	}

	return &newUser, nil
}
