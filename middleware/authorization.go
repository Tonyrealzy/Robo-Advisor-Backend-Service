package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JWTAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Log.Println("Authorization header missing")
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := GetUserClaims(db, tokenStr)
		if err != nil {
			logger.Log.Println("Invalid or expired token")
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if claims.ExpiresAt.Time.Before(time.Now().UTC()) {
			logger.Log.Println("Token has expired")
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "error": "Token has expired"})
			c.Abort()
			return
		}

		validUser, validErr := user.GetUserByID(db, claims.UserID)
		if validErr != nil {
			logger.Log.Printf("Error getting user By ID: %v", validErr)
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "error": validErr.Error()})
			c.Abort()
			return
		}

		c.Set("userClaims", claims)
		c.Set("user", validUser)
		c.Next()
	}
}
