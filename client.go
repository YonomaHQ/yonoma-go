package yonoma

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BaseURL = "https://api.yonoma.io/v1"

type Client struct {
	APIKey string
}

// NewClient initializes a new Yonoma API client
func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

// request sends an HTTP request and returns the response body
func (c *Client) request(method, endpoint string, payload interface{}) ([]byte, error) {
	url := BaseURL + endpoint
	var reqBody []byte
	var err error

	if payload != nil {
		reqBody, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	return body, nil
}
