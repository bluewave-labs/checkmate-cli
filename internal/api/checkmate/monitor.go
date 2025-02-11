package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bluewave-labs/checkmate-cli/internal/api/checkmate/types"
)

// CreateMonitor creates a new monitor
func (c *CheckmateClient) CreateMonitor(monitor *types.Monitor) (*http.Response, error) {
	validationErr := monitor.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	var requestBody []byte

	requestBody, err := json.Marshal(monitor)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	body := bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("POST", c.Credentials.APIBaseURL+"/monitors", body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Credentials.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.Credentials.APIKey)
	}

	return c.SendRequest(req)
}

// GetMonitor retrieves a monitor by ID
func (c *CheckmateClient) GetMonitor(id string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.Credentials.APIBaseURL+"/monitors/"+id, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Credentials.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.Credentials.APIKey)
	}

	return c.SendRequest(req)
}

// GetAllMonitors retrieves all monitors
func (c *CheckmateClient) GetAllMonitors() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.Credentials.APIBaseURL+"/monitors/", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Credentials.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.Credentials.APIKey)
	}

	return c.SendRequest(req)
}

// UpdateMonitor updates an existing monitor
func (c *CheckmateClient) UpdateMonitor(id string, monitor *types.Monitor) (*http.Response, error) {
	var requestBody []byte

	requestBody, err := json.Marshal(monitor)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	body := bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("PUT", c.Credentials.APIBaseURL+"/monitors/"+id, body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Credentials.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.Credentials.APIKey)
	}

	return c.SendRequest(req)
}

// DeleteMonitor deletes a monitor by ID
func (c *CheckmateClient) DeleteMonitor(id string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", c.Credentials.APIBaseURL+"/monitors/"+id, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Credentials.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.Credentials.APIKey)
	}

	return c.SendRequest(req)
}

func (c *CheckmateClient) CreateBulkMonitors(monitors []types.Monitor) (*http.Response, error) {
	// Validate each monitor
	for _, monitor := range monitors {
		if err := monitor.Validate(); err != nil {
			return nil, err
		}
	}

	var requestBody []byte

	requestBody, err := json.Marshal(monitors)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	body := bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("POST", c.Credentials.APIBaseURL+"/monitors/bulk", body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Credentials.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.Credentials.APIKey)
	}

	return c.SendRequest(req)
}
