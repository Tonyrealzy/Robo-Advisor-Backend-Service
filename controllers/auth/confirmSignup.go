package auth

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Confirm mail by link sent
// @Description  Confirm email used for signup
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      models.ConfirmSignupRequest  true  "Token and email for confirmation"
// @Success      200   {object}  models.ConfirmSignupResponse
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ServerErrorResponse
// @Router       /auth/signup/confirm [post]
func (base *Controller) ConfirmSignup(c *gin.Context) {
	var input models.ConfirmSignupRequest

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	message, messageErr := auth.ConfirmSignup(base.Db, input.Email, input.Token)
	if messageErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": messageErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
}
