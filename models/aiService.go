package models

import (
	"github.com/Tonyrealzy/go-backend/config"
	
	"time"
	"log"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AIServiceRequest struct {
	Age                 int     `json:"age"`
	Location            string  `json:"location"`
	InvestmentKnowledge string  `json:"investmentKnowledge"`
	InvestmentPurpose   string  `json:"investmentPurpose"`
	InvestmentHorizon   int     `json:"investmentHorizon"`
	RiskTolerance       string  `json:"riskTolerance"`
	Amount              float64 `json:"amount"`
	Currency            string  `json:"currency"`
}

type AIServiceResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type AIPersistedResponse struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	UserID    string         `json:"userId"`
	User      User           `gorm:"foreignKey:UserID"`
	Status    string         `json:"status"`
	Query     datatypes.JSON `gorm:"type:jsonb" json:"query"`
	Data      datatypes.JSON `gorm:"type:jsonb" json:"data,omitempty"`
	Message   string         `json:"message,omitempty"`
	Error     string         `json:"error,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

func (a *AIPersistedResponse) GetAllAIResponses(db *gorm.DB, userID string) (*AIPersistedResponse, error) {
	var interaction AIPersistedResponse

	err := config.FindByID(db, interaction, userID)
	if err != nil {
		log.Printf("Error getting all AI responses: %v", err)
		return nil, err
	}

	return &interaction, nil
}

func (a *AIPersistedResponse) GetTodayResponse(db *gorm.DB, userID string, pagination config.Pagination) ([]AIPersistedResponse, error) {
	today := time.Now()
	start := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	end := start.Add(24 * time.Hour)

	var results []AIPersistedResponse
	err := config.FindByThreeFieldsPaginated(db, results, "user_id", userID,
		"created_at >= ?", start,
		"created_at <= ?", end,
		pagination)
	if err != nil {
		log.Printf("Error getting today's AI responses: %v", err)
		return nil, err
	}

	return results, err
}

func (a *AIPersistedResponse) GetResponseByNoOfDays(db *gorm.DB, userID string, days int, pagination config.Pagination) ([]AIPersistedResponse, error) {
	now := time.Now()
	start := now.AddDate(0, 0, -days)

	var results []AIPersistedResponse
	err := config.FindByUserAndDateRangePaginated(db, &results, userID, start, now, pagination)
	if err != nil {
		log.Printf("Error getting all AI responses by number of days: %v", err)
		return nil, err
	}

	return results, err
}
