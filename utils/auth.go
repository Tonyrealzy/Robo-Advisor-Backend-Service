package utils

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/gofrs/uuid"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GenerateUUID() string {
	id, _ := uuid.NewV7()
	return id.String()
}

func IsValidUUID(id string) bool {
	_, err := uuid.FromString(id)
	return err == nil
}