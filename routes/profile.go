package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/controllers/profile"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/middleware"
)

func SetupProfileRoutes(router *gin.RouterGroup, controller profile.Controller) {
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware(controller.Db))

	protected.POST("profile", controller.GetProfile)

	health := router.Group("/")
	health.GET("health", controller.GetHealth)
}
