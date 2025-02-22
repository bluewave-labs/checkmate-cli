package types

import (
	"errors"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
)

type MonitorType string

// String enum for monitor types
const (
	MonitorHTTP            MonitorType = "http"
	MonitorPing            MonitorType = "ping"
	MonitorPageSpeed       MonitorType = "pagespeed"
	MonitorHardware        MonitorType = "hardware"
	MonitorDocker          MonitorType = "docker"
	MonitorPort            MonitorType = "port"
	MonitorDistributedHTTP MonitorType = "distributed_http"
)

var ErrForbidden = errors.New("403 forbidden. Please make sure you have the correct permissions.")
var ErrUnauthorized = errors.New("401 unauthorized. Please make sure you are logged in with the correct credentials.")
var ErrNotFound = errors.New("404 not found. Please make sure you are using the correct endpoint.")
var ErrBadRequest = errors.New("400 bad request.")
var ErrUnprocessableEntity = errors.New("422 unprocessable entity. Please make sure you are using the correct data format.")
var ErrClientError = errors.New("4xx client side error occured.")
var ErrServerError = errors.New("5xx server side error occured.")

// This struct represents a monitor object.
// Validations are done using the go-playground/validator package to ensure that the monitor data is valid.
// This validations should follow the same rules as the ones in the Checkmate API.
type Monitor struct {
	UserID   string      `json:"user_id" valid:"required"`
	TeamID   string      `json:"team_id" valid:"required"`
	Name     string      `json:"name" valid:"required,stringlength(1|100)"`
	URL      string      `json:"url" valid:"required,url"`
	Type     MonitorType `json:"type" valid:"in(http|ping|pagespeed|hardware|docker|port|distributed_http)"`
	IsActive bool        `json:"is_active"` // 'false' is for Paused monitors
	Status   bool        `json:"status"`    // 'false' is for Down and 'true' is for Up monitors
}

type MonitorTemplate struct {
	InstanceURL    string
	TeamID         string
	TotalMonitors  []string
	UpMonitors     []Monitor
	DownMonitors   []Monitor
	PausedMonitors []Monitor
}

func (m *MonitorTemplate) MonitorTable() [][]string {
	var result [][]string

	for _, monitor := range m.UpMonitors {
		result = append(result, []string{monitor.Name, monitor.URL, string(monitor.Type), color.GreenString("Up")})
	}

	for _, monitor := range m.DownMonitors {
		result = append(result, []string{monitor.Name, monitor.URL, string(monitor.Type), color.RedString("Down")})
	}

	for _, monitor := range m.PausedMonitors {
		result = append(result, []string{monitor.Name, monitor.URL, string(monitor.Type), color.YellowString("Paused")})
	}

	return result
}

func (m *Monitor) Validate() error {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return err
	}
	return nil
}
