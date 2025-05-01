package auth

import (
	"net/http"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services/auth"

	"github.com/gin-gonic/gin"
)

// @Summary      User logout
// @Description  Invalidate user session/token
// @Tags         Auth
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body  body      models.LogoutRequest  true  "Email for logout"
// @Success      200   {object}  models.LogoutResponse
// @Failure      401   {object}  models.AuthErrorResponse
// @Failure      500   {object}  models.ServerErrorResponse
// @Security BearerAuth
// @Router       /auth/logout [post]
func (base *Controller) Logout(c *gin.Context) {
	var input models.LogoutRequest

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	logoutErr := auth.Logout(base.Db, input.Email)
	if logoutErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": logoutErr.Error()})
		return
	}

	logger.Log.Println("Response successful!")
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Action successful"})
}
