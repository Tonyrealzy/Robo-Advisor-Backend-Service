package profile

import (
	"net/http"
	"time"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/gin-gonic/gin"
)

// @Summary      Health Check
// @Description  Returns service health status
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200   {object}  models.HealthResponse
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.AuthErrorResponse
// @Failure      500   {object}  models.ServerErrorResponse
// @Router       /health [get]
func (base *Controller) GetHealth(c *gin.Context) {
	logger.Log.Println("Response successful!")

	response := models.HealthResponse{
		Status:    "success",
		Message:   "Server is up and running",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	c.JSON(http.StatusOK, response)
}
