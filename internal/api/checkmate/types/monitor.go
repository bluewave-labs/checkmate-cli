package types

import "github.com/go-playground/validator/v10"

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

// This struct represents a monitor object.
// Validations are done using the go-playground/validator package to ensure that the monitor data is valid.
// This validations should follow the same rules as the ones in the Checkmate API.
type Monitor struct {
	UserID string      `json:"user_id" valid:"required"`
	TeamID string      `json:"team_id" valid:"required"`
	Name   string      `json:"name" valid:"required,stringlength(1|100)"`
	URL    string      `json:"url" valid:"required,url"`
	Type   MonitorType `json:"type" valid:"in(http|ping|pagespeed|hardware|docker|port|distributed_http)"`
}

func (m *Monitor) Validate() error {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return err
	}
	return nil
}
