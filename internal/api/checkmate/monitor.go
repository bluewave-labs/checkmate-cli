package checkmate

import (
	"bytes"
	"encoding/json"
	"io"
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
		return nil, err
	}

	body := bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("POST", c.credentials.APIBaseURL+"/monitors", body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.credentials.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.credentials.APIKey)
	}

	return c.SendRequest(req)
}

// GetMonitor retrieves a monitor by ID
func (c *CheckmateClient) GetMonitor(id string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.credentials.APIBaseURL+"/monitors/"+id, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.credentials.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.credentials.APIKey)
	}

	return c.SendRequest(req)
}

// GetAllMonitors retrieves all monitors
func (c *CheckmateClient) GetAllMonitors() (*types.APIResponse, error) {
	req, err := http.NewRequest("GET", c.credentials.APIBaseURL+"/monitors/", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.credentials.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.credentials.APIKey)
	}

	monitorResponse, err := c.SendRequest(req)
	if err != nil {
		return nil, err
	}
	switch monitorResponse.StatusCode {
	case http.StatusForbidden:
		return nil, types.ErrForbidden
	case http.StatusUnauthorized:
		return nil, types.ErrUnauthorized
	case http.StatusNotFound:
		return nil, types.ErrNotFound
	case http.StatusBadRequest:
		return nil, types.ErrBadRequest
	case http.StatusUnprocessableEntity:
		return nil, types.ErrUnprocessableEntity
	}

	defer monitorResponse.Body.Close()

	var apiResponse types.APIResponse
	// Read the response body
	body, err := io.ReadAll(monitorResponse.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	// Print the response body as a string
	return &apiResponse, nil
}

// UpdateMonitor updates an existing monitor
func (c *CheckmateClient) UpdateMonitor(id string, monitor *types.Monitor) (*http.Response, error) {
	var requestBody []byte

	requestBody, err := json.Marshal(monitor)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("PUT", c.credentials.APIBaseURL+"/monitors/"+id, body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.credentials.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.credentials.APIKey)
	}

	return c.SendRequest(req)
}

// DeleteMonitor deletes a monitor by ID
func (c *CheckmateClient) DeleteMonitor(id string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", c.credentials.APIBaseURL+"/monitors/"+id, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.credentials.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.credentials.APIKey)
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
		return nil, err
	}

	body := bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("POST", c.credentials.APIBaseURL+"/monitors/bulk", body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.credentials.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.credentials.APIKey)
	}

	return c.SendRequest(req)
}
