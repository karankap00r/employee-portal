package api

import (
	"encoding/json"
	"github.com/karankap00r/employee_portal/dto"
	"github.com/karankap00r/employee_portal/service"
	"github.com/karankap00r/employee_portal/storage/model"
	"github.com/karankap00r/employee_portal/util"
	"net/http"
	"time"
)

// RemoteWorkHandler is the interface for the remote work handler
type RemoteWorkHandler struct {
	service service.RemoteWorkService
}

// NewRemoteWorkHandler creates a new remote work handler with the given service
func NewRemoteWorkHandler(service service.RemoteWorkService) *RemoteWorkHandler {
	return &RemoteWorkHandler{service: service}
}

// GetRemoteWorkBalance gets the remote work balance for the employee with the given employee ID
func (h *RemoteWorkHandler) GetRemoteWorkBalance(w http.ResponseWriter, r *http.Request) {
	var request dto.GetRemoteWorkBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	balance, err := h.service.GetRemoteWorkBalance(r.Context(), request.EmployeeID)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response := dto.GetRemoteWorkBalanceResponse{
		EmployeeID:    balance.EmployeeID,
		Type:          balance.Type,
		AnnualBalance: balance.AnnualBalance,
	}
	util.WriteSuccessResponse(w, response)
}

// RaiseRemoteWorkRequest raises a remote work request with the given details
func (h *RemoteWorkHandler) RaiseRemoteWorkRequest(w http.ResponseWriter, r *http.Request) {
	var request dto.RaiseRemoteWorkRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	startDate, err := time.Parse(time.RFC3339, request.StartDate)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid start date format")
		return
	}
	endDate, err := time.Parse(time.RFC3339, request.EndDate)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid end date format")
		return
	}
	remoteWorkRequest := model.RemoteWorkRequest{
		EmployeeID: request.EmployeeID,
		Type:       request.Type,
		StartDate:  startDate,
		EndDate:    endDate,
		Reason:     request.Reason,
		Status:     "pending",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := h.service.RaiseRemoteWorkRequest(r.Context(), remoteWorkRequest); err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, "Remote work request raised successfully")
}

// UpdateRemoteWorkRequest updates the remote work request with the given details
func (h *RemoteWorkHandler) UpdateRemoteWorkRequest(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateRemoteWorkRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.UpdateRemoteWorkRequestStatus(r.Context(), request.RequestID, request.Status, request.UpdatedBy); err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, "Remote work request updated successfully")
}

// GetRemoteWorkRequestsInRange gets the remote work requests in the given date range
func (h *RemoteWorkHandler) GetRemoteWorkRequestsInRange(w http.ResponseWriter, r *http.Request) {
	var request dto.GetRemoteWorkRequestsInRangeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	startDate, err := time.Parse(time.RFC3339, request.StartDate)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid start date format")
		return
	}
	endDate, err := time.Parse(time.RFC3339, request.EndDate)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid end date format")
		return
	}
	remoteWorkRequests, err := h.service.GetRemoteWorkRequestsInRange(r.Context(), startDate, endDate)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	var response []dto.GetRemoteWorkRequestsInRangeResponse
	for _, remoteWorkRequest := range remoteWorkRequests {
		response = append(response, dto.GetRemoteWorkRequestsInRangeResponse{
			ID:         remoteWorkRequest.ID,
			EmployeeID: remoteWorkRequest.EmployeeID,
			Type:       remoteWorkRequest.Type,
			StartDate:  remoteWorkRequest.StartDate.Format(time.RFC3339),
			EndDate:    remoteWorkRequest.EndDate.Format(time.RFC3339),
			Status:     remoteWorkRequest.Status,
		})
	}
	util.WriteSuccessResponse(w, response)
}
