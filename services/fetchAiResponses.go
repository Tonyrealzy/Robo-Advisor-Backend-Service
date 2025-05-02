package services

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"

	"gorm.io/gorm"
)

func FetchAIResponsesForToday(db *gorm.DB, userID string, pagination config.Pagination) ([]models.AIServiceResponse, error) {
	var response *models.AIPersistedResponse

	persistedResponse, err := response.GetTodayResponse(db, userID, pagination)
	if err != nil {
		logger.Log.Printf("Error fetching AI response for today: %v", err)
		return nil, err
	}
	
	var transformedResponses []models.AIServiceResponse

	for _, r := range persistedResponse {
		transformed := models.AIServiceResponse{
			Status:  r.Status,
			Data:    r.Data,
			Message: r.Message,
			Error:   r.Error,
		}
		transformedResponses = append(transformedResponses, transformed)
	}

	return transformedResponses, nil
}

func FetchAIResponsesByNoOfDays(db *gorm.DB, userID string, number int, pagination config.Pagination) ([]models.AIServiceResponse, error) {
	var response models.AIPersistedResponse
	
	persistedResponse, err := response.GetResponseByNoOfDays(db, userID, number, pagination)
	if err != nil {
		logger.Log.Printf("Error fetching AI response by number of days: %v", err)
		return nil, err
	}

	var transformedResponses []models.AIServiceResponse

	for _, r := range persistedResponse {
		transformed := models.AIServiceResponse{
			Status:  r.Status,
			Data:    r.Data,
			Message: r.Message,
			Error:   r.Error,
		}
		transformedResponses = append(transformedResponses, transformed)
	}

	return transformedResponses, nil
}
