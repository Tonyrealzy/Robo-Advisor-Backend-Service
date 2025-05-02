package auth

import (
	"net/http"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services/auth"

	"github.com/gin-gonic/gin"
)

// @Summary      Resend verification link
// @Description  Resend verification link to the email used in signup/login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      models.PasswordResetRequest  true  "email"
// @Success      200   {object}  models.ResendLinkResponse
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ServerErrorResponse
// @Router       /auth/resend-link [post]
func (base *Controller) ResendLink(c *gin.Context) {
	var input models.PasswordResetRequest

	err := c.BindJSON(&input)
	if err != nil {
		logger.Log.Printf("Binding error: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	message, messageErr := auth.ResendLinkToUser(base.Db, input.Email)
	if messageErr != nil {
		logger.Log.Printf("ResendLinkToUser error: %v", messageErr.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": messageErr.Error()})
		return
	}

	logger.Log.Println("Response successful!")
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
}
