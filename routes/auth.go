package routes

import (
	"github.com/gin-gonic/gin"

	"go-backend/controllers/auth"
)

func SetupAuthRoutes(router *gin.RouterGroup, controller auth.Controller) {
	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)
	router.POST("/logout", controller.Logout)
	router.POST("/password-reset", controller.PasswordReset)
	router.POST("/change-password", controller.PasswordChange)
}
