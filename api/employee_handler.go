package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/karankap00r/employee_portal/dto/request"
	service "github.com/karankap00r/employee_portal/service/employee"
	"github.com/karankap00r/employee_portal/util"
	"net/http"
)

type EmployeeHandler interface {
	CreateEmployee(w http.ResponseWriter, r *http.Request)
	UpdateEmployee(w http.ResponseWriter, r *http.Request)
	GetEmployee(w http.ResponseWriter, r *http.Request)
	GetAllEmployees(w http.ResponseWriter, r *http.Request)
}

type employeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(service service.EmployeeService) EmployeeHandler {
	return &employeeHandler{service}
}

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
	employee, err := h.service.CreateEmployee(request)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, employee)
}

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
	employee, err := h.service.UpdateEmployeeByEmployeeID(employeeID, request)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, employee)
}

func (h *employeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	employee, err := h.service.GetEmployeeByEmployeeID(params["employeeID"])
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

func (h *employeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.service.GetAllEmployees()
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, employees)
}
