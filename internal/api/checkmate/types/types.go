package types

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}
