package ai

import (
	"net/http"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      AI Service
// @Description  Retrieval of previous responses from the AI Service
// @Tags         AI
// @Accept       json
// @Produce      json
// @Param        page   query     int     false  "Page number"
// @Param        limit  query     int     false  "Number of items per page"
// @Param        days   query     int     true   "Number of days to go back"
// @Success      200   {object}  models.AIResponse
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.AuthErrorResponse
// @Security BearerAuth
// @Router       /ai/fetch-response/by-days [get]
func (base *Controller) GetPreviousAiResponseByNoOfDays(c *gin.Context) {
	userRaw, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "error": "Invalid or expired token"})
		return
	}

	user, ok := userRaw.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Failed to fetch user details"})
		return
	}

	daysStr := c.Query("days")
	if daysStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "days parameter is required"})
		return
	}

	days, err := strconv.Atoi(daysStr)
	if err != nil || days <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "invalid days value"})
		return
	}

	pagination := config.GetPagination(c)
	responses, err := services.FetchAIResponsesByNoOfDays(base.Db, user.ID, days, pagination)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   responses,
	})
}
