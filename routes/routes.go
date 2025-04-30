package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"go-backend/controllers/auth"
	"go-backend/controllers/ai"
	_ "go-backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authController := auth.Controller{Db: db}
	aiController := ai.Controller{Db: db}

	authGroup := router.Group("/auth")
	SetupAuthRoutes(authGroup, authController)

	aiGroup := router.Group("/ai")
	SetupAIRoutes(aiGroup, aiController)
}
