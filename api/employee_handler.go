package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/karankap00r/employee_portal/dto/request"
	service "github.com/karankap00r/employee_portal/service/employee"
	"github.com/karankap00r/employee_portal/util"
)

type EmployeeHandler interface {
	GetEmployees(w http.ResponseWriter, r *http.Request)
	GetEmployee(w http.ResponseWriter, r *http.Request)
	CreateEmployee(w http.ResponseWriter, r *http.Request)
	UpdateEmployee(w http.ResponseWriter, r *http.Request)
}

type employeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(service service.EmployeeService) EmployeeHandler {
	return &employeeHandler{service}
}

func (h *employeeHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.service.GetAllEmployees()
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, employees)
}

func (h *employeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}
	employee, err := h.service.GetEmployeeByID(id)
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
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}
	var request request.UpdateEmployeeByEmployeeIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := request.Validate(); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	employee, err := h.service.UpdateEmployeeByID(id, request)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, employee)
}
