package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/controllers/auth"
)

func SetupAuthRoutes(router *gin.RouterGroup, controller auth.Controller) {
	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)
	router.POST("/logout", controller.Logout)
	router.POST("/password-reset", controller.PasswordReset)
	router.POST("/change-password", controller.PasswordChange)
}
