package profile

import (
	"net/http"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

// @Summary      User Profile
// @Description  Get user details by email
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        body  body      models.ProfileRequest  true  "Email"
// @Success      200   {object}  models.ProfileResponse
// @Failure      400   {object}  models.ErrorResponse
// @Failure      401   {object}  models.AuthErrorResponse
// @Failure      500   {object}  models.ServerErrorResponse
// @Router       /profile [post]
func (base *Controller) GetProfile(c *gin.Context) {
	var input models.ProfileRequest

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	userDetails, err := services.GetUserDetails(base.Db, input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	logger.Log.Println("Response successful!")
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": userDetails})
}
