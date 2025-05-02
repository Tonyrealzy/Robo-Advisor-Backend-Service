package middleware

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func CreateToken(userID string, email string) (string, error) {
	expirationMinutes, err := strconv.Atoi(config.AppConfig.JwtExpiration)
	if err != nil {
		logger.Log.Printf("JWT Token creation error: %v. Invalid JWT expiration value", err)
		return "", fmt.Errorf("invalid JWT expiration value: %v", err)
	}

	expirationDuration := time.Duration(expirationMinutes) * time.Minute

	claims := models.JWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expirationDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			Issuer:    "roboadvisor-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := config.AppConfig.JwtSecret

	return token.SignedString([]byte(secret))
}

func VerifyToken(db *gorm.DB, tokenStr string) (*models.JWTClaims, error) {
	parser := jwt.NewParser(jwt.WithoutClaimsValidation())

	token, err := parser.ParseWithClaims(tokenStr, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JwtSecret), nil
	})

	if err != nil {
		logger.Log.Printf("JWT error: %v", err)
		return nil, errors.New("invalid or expired token")
	}
	if !token.Valid {
		logger.Log.Println("Token is not valid")
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok {
		return nil, errors.New("could not parse claims")
	}

	var session models.UserSession
	_, sessionErr := session.GetUserSession(db, claims.UserID, tokenStr)
	if sessionErr != nil {
		logger.Log.Printf("Error getting user session from JWT session: %v", sessionErr)
		return nil, sessionErr
	}

	if time.Now().UTC().After(claims.ExpiresAt.Time) {
		logger.Log.Println("token has expired")
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
