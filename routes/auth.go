package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/controllers/auth"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/middleware"
)

func SetupAuthRoutes(router *gin.RouterGroup, controller auth.Controller) {
	router.POST("/signup", controller.Signup)
	router.POST("/signup/confirm", controller.ConfirmSignup)
	router.POST("/resend-link", controller.ResendLink)
	router.POST("/login", controller.Login)
	router.POST("/password-reset", controller.PasswordReset)
	router.POST("/change-password", controller.PasswordChange)
	
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware(controller.Db))
	protected.POST("/logout", controller.Logout)
}
