package services

import (
	"go-backend/config"
	"go-backend/models"

	"gorm.io/gorm"
)

func FetchAIResponsesForToday(db *gorm.DB, userID string, pagination config.Pagination) ([]models.AIServiceResponse, error) {
	var response *models.AIPersistedResponse

	persistedResponse, err := response.GetTodayResponse(db, userID, pagination)
	if err != nil {
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
	var response *models.AIPersistedResponse

	persistedResponse, err := response.GetResponseByNoOfDays(db, userID, number, pagination)
	if err != nil {
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
