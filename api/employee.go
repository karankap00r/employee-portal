package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Employee represents an employee
type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Dept string `json:"dept"`
}

// Employees in-memory list of employees
var Employees []Employee

// GetEmployees retrieves all employees
// @Summary Get all employees
// @Description Get all employees
// @Tags employees
// @Accept  json
// @Produce  json
// @Success 200 {array} Employee
// @Router /employees [get]
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Employees)
}

// GetEmployee retrieves a single employee by ID
// @Summary Get an employee by ID
// @Description Get an employee by ID
// @Tags employees
// @Accept  json
// @Produce  json
// @Param id path string true "Employee ID"
// @Success 200 {object} Employee
// @Failure 404 {string} string "Employee not found"
// @Router /employees/{id} [get]
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Employees {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}

// CreateEmployee creates a new employee
// @Summary Create a new employee
// @Description Create a new employee
// @Tags employees
// @Accept  json
// @Produce  json
// @Param employee body Employee true "Employee"
// @Success 201 {object} Employee
// @Router /employees [post]
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	Employees = append(Employees, employee)
	json.NewEncoder(w).Encode(employee)
}

// UpdateEmployee updates an existing employee by ID
// @Summary Update an employee
// @Description Update an employee
// @Tags employees
// @Accept  json
// @Produce  json
// @Param id path string true "Employee ID"
// @Param employee body Employee true "Employee"
// @Success 200 {object} Employee
// @Failure 404 {string} string "Employee not found"
// @Router /employees/{id} [put]
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Employees {
		if item.ID == params["id"] {
			Employees = append(Employees[:index], Employees[index+1:]...)
			var employee Employee
			_ = json.NewDecoder(r.Body).Decode(&employee)
			employee.ID = params["id"]
			Employees = append(Employees, employee)
			json.NewEncoder(w).Encode(employee)
			return
		}
	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}
