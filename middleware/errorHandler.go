package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"status":  "error",
					"message": "Internal Server Error",
					"details": fmt.Sprintf("%v", err),
				})
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			c.JSON(-1, gin.H{
				"status":  "error",
				"message": c.Errors[0].Error(),
			})
		}
	}
}
