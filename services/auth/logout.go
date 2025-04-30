package auth

import (
	"fmt"
	"go-backend/models"

	"gorm.io/gorm"
)

func Logout(db *gorm.DB, email string) error {
	var user *models.User
	var session *models.UserSession

	loggedInUser, err := user.GetUserByEmail(db, email)
	if err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	logoutErr := session.DeleteUserSession(db, loggedInUser.ID)
	if logoutErr != nil {
		return logoutErr
	}

	return nil
}
