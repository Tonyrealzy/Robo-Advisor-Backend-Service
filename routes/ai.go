package routes

import (
	"github.com/gin-gonic/gin"

	"robo-advisor-backend-service/controllers/ai"
)

func SetupAIRoutes(router *gin.RouterGroup, controller ai.Controller) {
	router.POST("/send-request", controller.GetAiResponse)
	router.GET("/fetch-response/today", controller.GetPreviousAiResponseForToday)
	router.GET("/fetch-response/by-days", controller.GetPreviousAiResponseByNoOfDays)
}
