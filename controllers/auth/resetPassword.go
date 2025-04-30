package auth

import (
	"net/http"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services/auth"

	"github.com/gin-gonic/gin"
)

// @Summary      Request password reset
// @Description  Generate password reset token and send email
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      models.PasswordResetRequest  true  "Email for password reset"
// @Success      200   {object}  models.PasswordResetResponse
// @Failure      400   {object}  models.ErrorResponse
// @Router       /auth/password-reset [post]
func (base *Controller) PasswordReset(c *gin.Context) {
	var input models.PasswordResetRequest

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	resetMsg, resetErr := auth.ResetPassword(base.Db, input.Email)
	if resetErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": resetErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": resetMsg})
}
