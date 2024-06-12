package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"

	"github.com/karankap00r/employee_portal/common"
	"github.com/karankap00r/employee_portal/dto"
	"github.com/karankap00r/employee_portal/service"
	"github.com/karankap00r/employee_portal/storage/model"
	"github.com/karankap00r/employee_portal/util"
)

// LeaveHandler is the interface for the leave handler
type LeaveHandler struct {
	service service.LeaveService
}

// NewLeaveHandler creates a new leave handler with the given service
func NewLeaveHandler(service service.LeaveService) *LeaveHandler {
	return &LeaveHandler{service: service}
}

// GetLeaveBalance gets the leave balance for the employee with the given employee ID
func (h *LeaveHandler) GetLeaveBalance(w http.ResponseWriter, r *http.Request) {
	var request dto.GetLeaveBalanceRequest

	params := mux.Vars(r)
	request.EmployeeID = params["employeeID"]

	balance, err := h.service.GetLeaveBalance(r.Context(), request.EmployeeID)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response := dto.GetLeaveBalanceResponse{
		EmployeeID:    balance.EmployeeID,
		LeaveType:     balance.LeaveType,
		AnnualBalance: balance.AnnualBalance,
	}
	util.WriteSuccessResponse(w, response)
}

// RaiseLeaveRequest raises a leave request with the given details
func (h *LeaveHandler) RaiseLeaveRequest(w http.ResponseWriter, r *http.Request) {
	var request dto.RaiseLeaveRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	startDate, err := time.Parse(time.DateOnly, request.StartDate)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid start date format")
		return
	}
	endDate, err := time.Parse(time.DateOnly, request.EndDate)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid end date format")
		return
	}
	leaveRequest := model.LeaveRequest{
		EmployeeID:    request.EmployeeID,
		LeaveCategory: request.LeaveCategory,
		StartDate:     startDate,
		EndDate:       endDate,
		Reason:        request.Reason,
		Status:        common.Pending.String(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	if err := h.service.RaiseLeaveRequest(r.Context(), leaveRequest); err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, "Leave request raised successfully")
}

// UpdateLeaveRequest updates the leave request with the given details
func (h *LeaveHandler) UpdateLeaveRequest(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateLeaveStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.UpdateLeave(r.Context(), request.LeaveID, request.Status, request.UpdatedBy); err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, "Leave request updated successfully")
}

// GetLeavesInRange gets the leave requests in the given date range
func (h *LeaveHandler) GetLeavesInRange(w http.ResponseWriter, r *http.Request) {
	var request dto.GetLeavesInRangeRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	startDate, err := time.Parse(time.DateOnly, request.StartDate)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid start date format")
		return
	}
	endDate, err := time.Parse(time.DateOnly, request.EndDate)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid end date format")
		return
	}
	leaveRequests, err := h.service.GetLeavesInRange(r.Context(), startDate, endDate)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	var response []dto.GetLeavesInRangeResponse
	for _, leaveRequest := range leaveRequests {
		response = append(response, dto.GetLeavesInRangeResponse{
			ID:            leaveRequest.ID,
			EmployeeID:    leaveRequest.EmployeeID,
			LeaveCategory: leaveRequest.LeaveCategory,
			StartDate:     leaveRequest.StartDate.Format(time.RFC3339),
			EndDate:       leaveRequest.EndDate.Format(time.RFC3339),
			Status:        leaveRequest.Status,
		})
	}
	util.WriteSuccessResponse(w, response)
}
