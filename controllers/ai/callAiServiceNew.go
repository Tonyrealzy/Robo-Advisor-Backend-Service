package ai

import (
	"net/http"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services"

	"github.com/gin-gonic/gin"
)

// @Summary      AI Service
// @Description  Interaction with the AI Service
// @Tags         AI
// @Accept       json
// @Produce      json
// @Param        body  body      models.AIServiceRequest  true  "Interaction with the Golang AI Service"
// @Success      200   {object}  models.AIResponse
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.AuthErrorResponse
// @Failure      500   {object}  models.ServerErrorResponse
// @Security BearerAuth
// @Router       /ai/request [post]
func (base *Controller) GetAIResponseNew(c *gin.Context) {
	var input models.AIServiceRequest

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

	err := c.BindJSON(&input)
	if err != nil {
		logger.Log.Println("Failed to bind JSON input")
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	resp, respErr := services.CallAIServiceNew(base.Db, base.Client, input, *user)
	if respErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": respErr.Error()})
		return
	}

	logger.Log.Println("Response successful!")
	c.JSON(http.StatusOK, resp)
}
