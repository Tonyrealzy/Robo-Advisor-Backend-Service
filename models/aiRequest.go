package models

import (
	"context"
	"fmt"

	genai "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIServiceImpl struct {
	client *genai.Client
}

func NewAIService(apiKey string) (*AIServiceImpl, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY is missing")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	return &AIServiceImpl{client: client}, nil
}

func (a *AIServiceImpl) GetAIResponse(message string) (string, error) {

	ctx := context.Background()
	model := a.client.GenerativeModel("models/gemini-1.5-flash-latest")

	resp, err := model.GenerateContent(ctx, genai.Text(message))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response from Gemini")
	}

	// Convert first response part to text
	if textPart, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
		return string(textPart), nil
	}

	return "", fmt.Errorf("unexpected response format")
}

func (a *AIServiceImpl) FineTunedResponse(req AIServiceRequest) (string, error) {
	startPrompt := fmt.Sprintf(
	`I am %v years old and I live in %s. I consider myself a %s in terms of investment and I wish to invest for the purpose of %s over a %v-year horizon, and have a %s risk tolerance.
	I have a sum of %v in %s to invest.

	Please recommend at least five specific financial products (including ticker and provider) that a typical financial advisor might suggest based on my profile.

	🔸 Return the response strictly in JSON format.
	🔸 The top-level key must be "recommendations".
	🔸 Each recommendation must contain the following keys in order:
	- "financial_product": string
	- "ticker": string
	- "provider": string
	- "brief_description": string
	- "expected_return": a percentage as a string (e.g., "7%%")
	- "principal": an integer amount from my total capital (e.g., 5000)
	- "estimated_return_value": an integer computed as (expected_return × principal). For example, if expected_return is "7%%" and principal is 1000, then estimated_return_value should be 70.
	- "composition": an integer percentage (adds up to 100 across all recommendations)

	🔸 Ensure the "composition" field is the last field.

	Important rules:
	- All numeric fields (principal, estimated_return_value, composition) must be integers only — no ranges, symbols, or units.
	- Do not include strings like "over 10 years" or "USD".
	- The total "composition" across recommendations must equal 100.

	I will not consider this personalized advice.`,
	req.Age, req.Location, req.InvestmentKnowledge, req.InvestmentPurpose, req.InvestmentHorizon, req.RiskTolerance, req.Amount, req.Currency)

	// startPrompt := fmt.Sprintf("I am %v years old and I live in %s. I consider myself a %s in terms of investment and I wish to invest for the purpose of %s over a %v-year horizon, and have a %s risk tolerance. I have a sum of %v in %s to invest. Which specific financial products (including ticker and provider) would a typical financial advisor recommend for investment given my circumstances? You may limit the number of recommendations to a minimum of five. You may also add an estimate of the expected return from each recommendation. Environmental factors are not important to me when I am investing. Which composition (as a percentage) would he recommend for each financial product? Each recommendation should be a proportion of '100%%' of the amount to be invested. I will not consider your response personalized advice. You may send the response in JSON format, let the key to the recommendations be given as recommendations, and each key in the recommendation should be represented as (`financial_product`, `ticker`, `provider`, `brief_description`, `expected_return`, `principal`, `estimated_return_value`, `composition`) and also ensure that composition comes last in the json.",
	// 	req.Age, req.Location, req.InvestmentKnowledge, req.InvestmentPurpose, req.InvestmentHorizon, req.RiskTolerance, req.Amount, req.Currency)

	ai_response, err := a.GetAIResponse(startPrompt)
	if err != nil {
		return "", err
	}

	return ai_response, err
}
