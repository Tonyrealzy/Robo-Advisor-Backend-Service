package services

import (
	"fmt"
	"gorm.io/gorm"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
)

func GetUserDetails(db *gorm.DB, email string) (*models.Profile, error) {
	var user models.User

	existingUser, err := user.GetUserByEmail(db, email)
	if err != nil {
		logger.Log.Printf("user not found: %v", err)
		return nil, fmt.Errorf("user not found: %v", err)
	}

	userDetails := models.Profile{
		Name:      existingUser.Name,
		Email:     existingUser.Email,
		FirstName: existingUser.FirstName,
		LastName:  existingUser.LastName,
		IsActive:  existingUser.IsActive,
	}

	return &userDetails, nil
}
