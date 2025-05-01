package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/controllers/ai"
)

func SetupAIRoutes(router *gin.RouterGroup, controller ai.Controller) {
	router.POST("/send-request", controller.GetAiResponse)
	router.GET("/fetch-response/today", controller.GetPreviousAiResponseForToday)
	router.GET("/fetch-response/days", controller.GetPreviousAiResponseByNoOfDays)
}
