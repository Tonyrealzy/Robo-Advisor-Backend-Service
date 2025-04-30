package middleware

import (
	"github.com/gin-gonic/gin"
)

func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prevent clickjacking
		c.Header("X-Frame-Options", "DENY")

		// Block XSS attacks
		c.Header("X-XSS-Protection", "1; mode=block")

		// Prevent MIME-sniffing
		c.Header("X-Content-Type-Options", "nosniff")

		// Enforce HTTPS (adjust max-age)
		c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		// Cross-Origin Resource Policy
		c.Header("Cross-Origin-Resource-Policy", "same-origin")

		c.Next()
	}
}
