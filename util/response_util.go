package util

import (
	"encoding/json"
	"net/http"

	"github.com/karankap00r/employee_portal/dto/response"
)

func WriteSuccessResponse(w http.ResponseWriter, data interface{}) {
	response := response.APIResponse{
		Success: true,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	response := response.APIResponse{
		Success: false,
		Error: &response.APIError{
			Code:    code,
			Message: message,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
