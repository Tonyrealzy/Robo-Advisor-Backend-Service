package auth

import (
	"net/http"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services/auth"

	"github.com/gin-gonic/gin"
)

// @Summary      Confirm email
// @Description  Use hashed token to confirm email used in signup
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
		logger.Log.Printf("Binding error: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	message, messageErr := auth.ConfirmSignup(base.Db, input.Email, input.Token)
	if messageErr != nil {
		logger.Log.Printf("ConfirmSignup error: %v", messageErr.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": messageErr.Error()})
		return
	}

	logger.Log.Println("Response successful!")
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
}
