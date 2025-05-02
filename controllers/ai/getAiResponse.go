package ai

import (
	"net/http"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services"

	"github.com/gin-gonic/gin"
)

// @Summary      AI Service
// @Description  Retrieval of previous responses from the AI Service
// @Tags         AI
// @Accept       json
// @Produce      json
// @Param        page   query     int     false  "Page number"
// @Param        limit  query     int     false  "Number of items per page"
// @Success      200   {object}  models.AIResponse
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.AuthErrorResponse
// @Failure      500   {object}  models.ServerErrorResponse
// @Security BearerAuth
// @Router       /ai/fetch-response/today  [get]
func (base *Controller) GetPreviousAiResponseForToday(c *gin.Context) {
	userRaw, exists := c.Get("user")
	if !exists {
		logger.Log.Println("Invalid or expired token")
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "error": "Invalid or expired token"})
		return
	}

	user, ok := userRaw.(*models.User)
	if !ok {
		logger.Log.Println("Failed to fetch user details")
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Failed to fetch user details"})
		return
	}

	pagination := config.GetPagination(c)
	responses, err := services.FetchAIResponsesForToday(base.Db, user.ID, pagination)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	logger.Log.Println("Response successful!")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   responses,
	})
}
