package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"github.com/Tonyrealzy/Robo-Advisor-Backend-Service/models"
)

func FormatResponse(rawText string) ([]models.Recommendation, error) {
	replacements := []struct {
		old string
		new string
	}{
		{"\n", ""},
		{"```", ""},
		{"json", ""},
		{"    ", ""},
		{"  ", ""},
	}
	for _, rep := range replacements {
		rawText = strings.ReplaceAll(rawText, rep.old, rep.new)
	}

	re := regexp.MustCompile(`\{.*"recommendations"\s*:\s*\[.*?\]\s*\}`)
	jsonPart := re.FindString(rawText)
	if jsonPart == "" {
		return nil, fmt.Errorf("could not extract JSON from response")
	}

	var parsed models.AIFirstResponse
	err := json.Unmarshal([]byte(jsonPart), &parsed)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}


	return parsed.Recommendations, nil
}
