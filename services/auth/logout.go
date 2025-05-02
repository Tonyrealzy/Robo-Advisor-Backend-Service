package auth

import (
	"fmt"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"

	"gorm.io/gorm"
)

func Logout(db *gorm.DB, email string) error {
	var user models.User
	var session models.UserSession

	loggedInUser, err := user.GetUserByEmail(db, email)
	if err != nil {
		logger.Log.Printf("user not found: %v", err)
		return fmt.Errorf("user not found: %v", err)
	}

	logoutErr := session.DeleteUserSession(db, loggedInUser.ID)
	if logoutErr != nil {
		logger.Log.Printf("Error deleting user session: %v", logoutErr)
		return logoutErr
	}

	return nil
}
