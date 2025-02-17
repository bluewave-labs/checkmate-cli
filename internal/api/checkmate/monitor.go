package checkmate

import (
	"github.com/bluewave-labs/checkmate-cli/internal/api/checkmate/types"
)

// CreateMonitor creates a new monitor
func (c *CheckmateClient) CreateMonitor(monitor *types.Monitor) (*types.APIResponse, error) {
	if err := monitor.Validate(); err != nil {
		return nil, err
	}

	return c.performRequest("POST", "/monitors", monitor, c.authenticator)
}

// GetMonitor retrieves a monitor by ID
func (c *CheckmateClient) GetMonitor(id string) (*types.APIResponse, error) {
	return c.performRequest("GET", "/monitors/"+id, nil, c.authenticator)
}

// GetAllMonitors retrieves all monitors
func (c *CheckmateClient) GetAllMonitors() (*types.APIResponse, error) {
	return c.performRequest("GET", "/monitors", nil, c.authenticator)
}

// UpdateMonitor updates an existing monitor
func (c *CheckmateClient) UpdateMonitor(id string, monitor *types.Monitor) (*types.APIResponse, error) {
	if err := monitor.Validate(); err != nil {
		return nil, err
	}

	return c.performRequest("PUT", "/monitors/"+id, monitor, c.authenticator)
}

// DeleteMonitor deletes a monitor by ID
func (c *CheckmateClient) DeleteMonitor(id string) (*types.APIResponse, error) {
	return c.performRequest("DELETE", "/monitors/"+id, nil, c.authenticator)
}

func (c *CheckmateClient) CreateBulkMonitors(monitors []types.Monitor) (*types.APIResponse, error) {
	// Validate each monitor
	for _, monitor := range monitors {
		if err := monitor.Validate(); err != nil {
			return nil, err
		}
	}

	return c.performRequest("POST", "/monitors/bulk", monitors, c.authenticator)
}
