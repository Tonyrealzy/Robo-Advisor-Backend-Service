package auth

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Register a new user
// @Description  Create a new user account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      models.SignupRequest  true  "Signup details"
// @Success      200   {object}  models.SignupResponse
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ServerErrorResponse
// @Router       /auth/signup [post]
func (base *Controller) Signup(c *gin.Context) {
	var input models.SignupRequest

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	_, linkMsg, signupErr := auth.Signup(base.Db, input.Email, input.Password, input.FirstName, input.LastName, input.Name)
	if signupErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": signupErr.Error()})
		return
	}

	logger.Log.Println("Response successful!")
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User created successfully", "token": linkMsg})
}
