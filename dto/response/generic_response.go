package response

// APIResponse represents the response of an API
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

// APIError represents the error in an API response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
