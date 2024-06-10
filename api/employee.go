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

// in-memory list of Employees
var Employees []Employee

// Get all Employees
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Employees)
}

// Get single employee
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

// Create a new employee
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	Employees = append(Employees, employee)
	json.NewEncoder(w).Encode(employee)
}

// Update an employee
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

// Delete an employee
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Employees {
		if item.ID == params["id"] {
			Employees = append(Employees[:index], Employees[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Employees)
}
