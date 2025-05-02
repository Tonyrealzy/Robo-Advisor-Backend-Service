package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash := sha256.Sum256([]byte(password))
	shaPass := hex.EncodeToString(hash[:])

	// Then hash with bcrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(shaPass), bcrypt.DefaultCost)
	return string(bytes), err
	// bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	// return string(bytes), err
}

func CheckPasswordHash(password, hashed string) bool {
	hash := sha256.Sum256([]byte(password))
	shaPass := hex.EncodeToString(hash[:])
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(shaPass))
	return err == nil
	// return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
}

func GenerateUUID() string {
	id, _ := uuid.NewV7()
	return id.String()
}

func IsValidUUID(id string) bool {
	_, err := uuid.FromString(id)
	return err == nil
}
