package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/utils"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func CallAIService(db *gorm.DB, req models.AIServiceRequest, user models.User) (*models.AIServiceResponse, error) {
	httpClient := utils.NewHTTPClient(180 * time.Second)
	aiServiceURL := fmt.Sprintf("%s/api/gemini_request", config.AppConfig.AiService)

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
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}

		createErr := config.CreateOneRecord(db, &aiServiceResponse)
		if createErr != nil {
			logger.Log.Printf("error creating AI Service chat entry: %v", createErr)
			return nil, createErr
		}
	}

	return &parsedResponse, nil
}

func PostToAI(client *models.AIServiceImpl, req models.AIServiceRequest) ([]models.Recommendation, error) {
	response, err := client.FineTunedResponse(req)
	if err != nil {
		return nil, err
	}

	formattedResp, err := utils.FormatResponse(response)
	if err != nil {
		return nil, err
	}

	return formattedResp, nil
}

func CallAIServiceNew(db *gorm.DB, client *models.AIServiceImpl, req models.AIServiceRequest, user models.User) (*models.AIServiceResponse, error) {

	response, err := PostToAI(client, req)
	if err != nil {
		logger.Log.Printf("failed to get response from AI: %v", err)
		return nil, fmt.Errorf("failed to get response from AI: %v", err)
	}

	queryJSON, err := json.Marshal(req)
	if err != nil {
		logger.Log.Printf("failed to marshal request: %v", err)
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	dataJSON, err := json.Marshal(response)
	if err != nil {
		logger.Log.Printf("failed to marshal response data: %v", err)
		return nil, fmt.Errorf("failed to marshal response data: %v", err)
	}

	aiServiceResponse := models.AIPersistedResponse{
		User:      user,
		UserID:    user.ID,
		Query:     datatypes.JSON(queryJSON),
		Data:      datatypes.JSON(dataJSON),
		Status:    "success",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	createErr := config.CreateOneRecord(db, &aiServiceResponse)
	if createErr != nil {
		logger.Log.Printf("error creating AI Service chat entry: %v", createErr)
		return nil, createErr
	}

	finalResp := models.AIServiceResponse{
		Status:  "success",
		Data:    datatypes.JSON(dataJSON),
	}

	return &finalResp, err
}
