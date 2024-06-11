package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karankap00r/employee_portal/dto/request"
	"github.com/karankap00r/employee_portal/service"
	"github.com/karankap00r/employee_portal/util"
)

// EmployeeHandler is the interface for the employee handler
type EmployeeHandler interface {

	// CreateEmployee creates the employee with the given details
	CreateEmployee(w http.ResponseWriter, r *http.Request)

	// UpdateEmployee updates the employee with the given details
	UpdateEmployee(w http.ResponseWriter, r *http.Request)

	// GetEmployee gets the employee with the given employee ID
	GetEmployee(w http.ResponseWriter, r *http.Request)

	// GetAllEmployees gets all the employees for the organization
	GetAllEmployees(w http.ResponseWriter, r *http.Request)
}

type employeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(service service.EmployeeService) EmployeeHandler {
	return &employeeHandler{service}
}

// CreateEmployee creates the employee with the given details
func (h *employeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var request request.CreateEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := request.Validate(); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	employee, err := h.service.CreateEmployee(r.Context(), request)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, employee)
}

// UpdateEmployee updates the employee with the given details
func (h *employeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	employeeID := params["employeeID"]
	var request request.UpdateEmployeeByEmployeeIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := request.Validate(); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	employee, err := h.service.UpdateEmployeeByEmployeeID(r.Context(), employeeID, request)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, employee)
}

// GetEmployee gets the employee with the given employee ID
func (h *employeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	employee, err := h.service.GetEmployeeByEmployeeID(r.Context(), params["employeeID"])
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if employee == nil {
		util.WriteErrorResponse(w, http.StatusNotFound, "Employee not found")
		return
	}
	util.WriteSuccessResponse(w, employee)
}

// GetAllEmployees gets all the employees for the organization
func (h *employeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.service.GetAllEmployees(r.Context())
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, employees)
}
