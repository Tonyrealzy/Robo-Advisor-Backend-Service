package services

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
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

	log.Printf("AI Service URL: %v", aiServiceURL)

	resp, body, err := httpClient.PostRequest(aiServiceURL, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call ai-service: %v", err)
	}

	log.Printf("Response Status Code: %v", resp.StatusCode)
	log.Printf("Response Headers: %v", resp.Header)

	var parsedResponse models.AIServiceResponse
	if err := json.Unmarshal(body, &parsedResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	if parsedResponse.Status == "success" {
		queryJSON, err := json.Marshal(req)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request: %v", err)
		}

		dataJSON, err := json.Marshal(parsedResponse.Data)
		if err != nil {
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
			return nil, createErr
		}
	}

	return &parsedResponse, nil
}
