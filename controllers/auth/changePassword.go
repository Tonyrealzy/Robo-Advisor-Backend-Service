package auth

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Request password reset
// @Description  Generate password reset token and send email
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      models.PasswordChangeRequest  true  "Token and new password for reset"
// @Success      200   {object}  models.PasswordChangeResponse
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ServerErrorResponse
// @Router       /auth/change-password [post]
func (base *Controller) PasswordChange(c *gin.Context) {
	var input models.PasswordChangeRequest

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	resetMsg, resetErr := auth.ValidateResetToken(base.Db, input.Token, input.NewPassword)
	if resetErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": resetErr.Error()})
		return
	}

	logger.Log.Println("Response successful!")
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": resetMsg})
}
