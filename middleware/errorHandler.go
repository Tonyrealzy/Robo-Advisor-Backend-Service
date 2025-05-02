package middleware

import (
	"fmt"
	"net/http"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Printf("Internal Server Error: %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"status":  "error",
					"message": "Internal Server Error",
					"details": fmt.Sprintf("%v", err),
				})
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			logger.Log.Printf("Internal Server Error: %v", c.Errors)
			c.JSON(-1, gin.H{
				"status":  "error",
				"message": c.Errors[0].Error(),
			})
		}
	}
}
