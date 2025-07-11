package models

import (
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/config"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/internal/logger"

	"encoding/json"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type IntString int

func (i *IntString) UnmarshalJSON(data []byte) error {
	var tempInt int
	if err := json.Unmarshal(data, &tempInt); err == nil {
		*i = IntString(tempInt)
		return nil
	}

	var tempStr string
	if err := json.Unmarshal(data, &tempStr); err == nil {
		parsedInt, err := strconv.Atoi(tempStr)
		if err != nil {
			return fmt.Errorf("cannot parse string to int: %v", err)
		}
		*i = IntString(parsedInt)
		return nil
	}

	return fmt.Errorf("unsupported type for IntString: %s", string(data))
}

type AIServiceRequest struct {
	Age                 int     `json:"age" validate:"required"`
	Location            string  `json:"location" validate:"required"`
	InvestmentKnowledge string  `json:"investmentKnowledge" validate:"required"`
	InvestmentPurpose   string  `json:"investmentPurpose" validate:"required"`
	InvestmentHorizon   int     `json:"investmentHorizon" validate:"required"`
	RiskTolerance       string  `json:"riskTolerance" validate:"required"`
	Amount              float64 `json:"amount" validate:"required"`
	Currency            string  `json:"currency" validate:"required"`
}

type Recommendation struct {
	FinancialProduct     string    `json:"financial_product"`
	Ticker               string    `json:"ticker"`
	Provider             string    `json:"provider"`
	BriefDescription     string    `json:"brief_description"`
	ExpectedReturn       string    `json:"expected_return"`
	Composition          IntString    `json:"composition"`
	Principal            IntString `json:"principal"`
	EstimatedReturnValue IntString `json:"estimated_return_value"`
}

type AIFirstResponse struct {
	Recommendations []Recommendation `json:"recommendations"`
}

type AIServiceResponse struct {
	Status    string      `json:"status"`
	Query     interface{} `json:"query,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Message   string      `json:"message,omitempty"`
	Error     string      `json:"error,omitempty"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
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

	err := config.FindByID(db, &interaction, userID)
	if err != nil {
		logger.Log.Printf("Error getting all AI responses: %v", err)
		return nil, err
	}

	return &interaction, nil
}

func (a *AIPersistedResponse) GetTodayResponse(db *gorm.DB, userID string, pagination config.Pagination) ([]AIPersistedResponse, error) {
	today := time.Now().UTC()
	start := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	end := start.Add(24 * time.Hour)

	var results []AIPersistedResponse
	err := config.FindByThreeFieldsPaginated(db, &results, "user_id = ?", userID,
		"created_at >= ?", start,
		"created_at <= ?", end,
		pagination)
	if err != nil {
		logger.Log.Printf("Error getting today's AI responses: %v", err)
		return nil, err
	}

	return results, err
}

func (a *AIPersistedResponse) GetResponseByNoOfDays(db *gorm.DB, userID string, days int, pagination config.Pagination) ([]AIPersistedResponse, error) {
	now := time.Now().UTC()
	start := now.AddDate(0, 0, -days)

	var results []AIPersistedResponse
	err := config.FindByUserAndDateRangePaginated(db, &results, userID, start, now, pagination)
	if err != nil {
		logger.Log.Printf("Error getting all AI responses by number of days: %v", err)
		return nil, err
	}

	return results, err
}

func (a *AIPersistedResponse) GetResponseByDateRange(db *gorm.DB, userID string, from, to time.Time, pagination config.Pagination) ([]AIPersistedResponse, error) {
	var results []AIPersistedResponse

	err := config.FindByUserAndDateRangePaginated(db, &results, userID, from, to, pagination)
	if err != nil {
		logger.Log.Printf("Error getting all AI responses by date range: %v", err)
		return nil, err
	}

	return results, err
}
