package routes

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/controllers/ai"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/controllers/auth"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/controllers/profile"
	_ "github.com/Tonyrealzy/Robo-Advisor-Backend-Service/docs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authController := auth.Controller{Db: db}
	aiController := ai.Controller{Db: db}
	profileController := profile.Controller{Db: db}

	authGroup := router.Group("/auth")
	SetupAuthRoutes(authGroup, authController)

	aiGroup := router.Group("/ai")
	SetupAIRoutes(aiGroup, aiController)

	profileGroup := router.Group("/")
	SetupProfileRoutes(profileGroup, profileController)
}
