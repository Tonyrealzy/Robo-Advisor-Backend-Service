package models

import (
	"time"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:uuid;primaryKey;unique;not null"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	IsActive  bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type SignupRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"username" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ProfileRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type LogoutRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (u *User) CreateUser(db *gorm.DB, user *User) error {
	err := config.CreateOneRecord(db, user)
	if err != nil {
		logger.Log.Printf("Error creating user: %v", err)
		return err
	}

	return nil
}

func (u *User) GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User

	err := config.FindOneByField(db, &user, "email", email)
	if err != nil {
		logger.Log.Printf("Error finding by one field: %v", err)
		return nil, err
	}

	return &user, nil
}

func (u *User) GetUserActivityByEmail(db *gorm.DB, email string) (*User, error) {
	var user User

	err := config.FindByTwoFields(db, &user, "email = ?", email, "is_active = ?", true)
	if err != nil {
		logger.Log.Printf("Error finding by email and activeness: %v", err)
		return nil, err
	}

	return &user, nil
}

func (u *User) IsUserActive(db *gorm.DB, isActive bool) (bool, error) {
	var user User

	err := config.FindOneByField(db, &user, "is_active", isActive)
	if err != nil {
		logger.Log.Printf("Error finding by one field: %v", err)
		return false, err
	}

	return user.IsActive, nil
}

func (u *User) GetUserByUsername(db *gorm.DB, name string) (*User, error) {
	var user User

	err := config.FindOneByField(db, &user, "name", name)
	if err != nil {
		logger.Log.Printf("Error getting user by username: %v", err)
		return nil, err
	}

	return &user, nil
}

func (u *User) GetUserByID(db *gorm.DB, id string) (*User, error) {
	var user User

	err := config.FindByID(db, &user, id)
	if err != nil {
		logger.Log.Printf("Error getting user by ID: %v", err)
		return nil, err
	}

	return &user, nil
}

func (u *User) DeleteUser(db *gorm.DB, userID string) error {
	err := config.HardDeleteSpecificRecord(db, u, "id = ?", userID)
	if err != nil {
		logger.Log.Printf("Error deleting user: %v", err)
		return err
	}

	return nil
}

func (u *User) UpdateUserPassword(db *gorm.DB, user *User) error {
	err := config.UpdateOneFieldByID(db, user, user.ID, "password", user.Password)
	if err != nil {
		logger.Log.Printf("Error updating user password: %v", err)
		return err
	}

	return nil
}
