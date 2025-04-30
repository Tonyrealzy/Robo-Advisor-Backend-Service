package middleware

import (
	"net/http"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JWTAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := GetUserClaims(db, tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if claims.ExpiresAt.Time.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "error": "Token has expired"})
			c.Abort()
			return
		}

		validUser, validErr := user.GetUserByID(db, claims.UserID)
		if validErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "error": validErr.Error()})
			c.Abort()
			return
		}

		c.Set("userClaims", claims)
		c.Set("user", validUser)
		c.Next()
	}
}
