package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	Client *http.Client
}

func NewHTTPClient(timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		Client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (h *HTTPClient) GetRequest(url string) (*http.Response, []byte, error) {
	resp, err := h.Client.Get(url)
	if err != nil {
		log.Printf("Failed to send GET request: %v", err)
		return nil, nil, fmt.Errorf("failed to send GET request: %v", err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return resp, body, nil
}

func (h *HTTPClient) PostRequest(url string, data interface{}) (*http.Response, []byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal data: %v", err)
		return nil, nil, fmt.Errorf("failed to marshal data: %v", err)
	}
	
	resp, err := h.Client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Failed to send POST request: %v", err)
		return nil, nil, fmt.Errorf("failed to send POST request: %v", err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return resp, body, nil
}
