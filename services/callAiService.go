package services

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/utils"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func CallAIService(db *gorm.DB, req models.AIServiceRequest, user models.User) (*models.AIServiceResponse, error) {
	httpClient := utils.NewHTTPClient(10 * time.Second)
	aiServiceURL := fmt.Sprintf("%s/api/gemini_request", config.AppConfig.AiService)

	logger.Log.Printf("AI Service URL: %v", aiServiceURL)
	log.Printf("AI Service URL: %v", aiServiceURL)
	
	resp, body, err := httpClient.PostRequest(aiServiceURL, req)
	if err != nil {
		logger.Log.Printf("failed to call ai-service: %v", err)
		return nil, fmt.Errorf("failed to call ai-service: %v", err)
	}
	
	logger.Log.Printf("Response Status Code: %v", resp.StatusCode)
	
	var parsedResponse models.AIServiceResponse
	if err := json.Unmarshal(body, &parsedResponse); err != nil {
		logger.Log.Printf("failed to parse JSON response: %v", err)
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}
	
	if parsedResponse.Status == "success" {
		queryJSON, err := json.Marshal(req)
		if err != nil {
			logger.Log.Printf("failed to marshal request: %v", err)
			return nil, fmt.Errorf("failed to marshal request: %v", err)
		}
		
		dataJSON, err := json.Marshal(parsedResponse.Data)
		if err != nil {
			logger.Log.Printf("failed to marshal response data: %v", err)
			return nil, fmt.Errorf("failed to marshal response data: %v", err)
		}

		aiServiceResponse := models.AIPersistedResponse{
			User:      user,
			UserID:    user.ID,
			Query:     datatypes.JSON(queryJSON),
			Data:      datatypes.JSON(dataJSON),
			Status:    parsedResponse.Status,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		createErr := config.CreateOneRecord(db, aiServiceResponse)
		if createErr != nil {
			logger.Log.Printf("error creating AI Service chat entry: %v", createErr)
			return nil, createErr
		}
	}

	return &parsedResponse, nil
}
