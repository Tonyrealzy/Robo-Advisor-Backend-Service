package middleware

import (
	"errors"
	"fmt"
	"github.com/Tonyrealzy/go-backend/config"
	"github.com/Tonyrealzy/go-backend/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func CreateToken(userID string, email string) (string, error) {
	expirationMinutes, err := strconv.Atoi(config.AppConfig.JwtExpiration)
	if err != nil {
		return "", fmt.Errorf("invalid JWT expiration value: %v", err)
	}

	expirationDuration := time.Duration(expirationMinutes) * time.Minute

	claims := models.JWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "roboadvisor-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := config.AppConfig.JwtSecret

	return token.SignedString([]byte(secret))
}

func VerifyToken(db *gorm.DB, tokenStr string) (*models.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok {
		return nil, errors.New("could not parse claims")
	}

	var session models.UserSession
	userSession, err := session.GetUserSession(db, claims.UserID, tokenStr)
	if err != nil {
		return nil, err
	}

	if time.Now().After(userSession.ExpiresAt) {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}

func IsTokenValid(db *gorm.DB, tokenStr string) bool {
	_, err := VerifyToken(db, tokenStr)
	return err == nil
}

func GetUserClaims(db *gorm.DB, tokenStr string) (*models.JWTClaims, error) {
	return VerifyToken(db, tokenStr)
}
