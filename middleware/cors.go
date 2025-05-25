package middleware

import (
	"net/http"
	"time"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://robo-advisor-frontend.netlify.app", "http://localhost:3000", "http://localhost:3001", "http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(ErrorHandlingMiddleware())
	r.Use(SecurityHeadersMiddleware())

	r.NoRoute(func(c *gin.Context) {
		logger.Log.Println("Route Not found")
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Route not found",
		})
	})

	return r
}
