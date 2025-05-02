package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/controllers/ai"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/middleware"
)

func SetupAIRoutes(router *gin.RouterGroup, controller ai.Controller) {
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware(controller.Db))

	protected.POST("/send-request", controller.GetAiResponse)
	protected.GET("/fetch-response/today", controller.GetPreviousAiResponseForToday)
	protected.GET("/fetch-response/days", controller.GetPreviousAiResponseByNoOfDays)
}
