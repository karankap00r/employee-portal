package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karankap00r/employee_portal/database"
)

// Employee represents an employee
type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Dept string `json:"dept"`
}

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
	rows, err := database.GetDB().Query("SELECT id, name, age, dept FROM employees")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var employee Employee
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Age, &employee.Dept)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employees = append(employees, employee)
	}
	json.NewEncoder(w).Encode(employees)
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
	var employee Employee
	err := database.GetDB().QueryRow("SELECT id, name, age, dept FROM employees WHERE id = ?", params["id"]).Scan(&employee.ID, &employee.Name, &employee.Age, &employee.Dept)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Employee not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(employee)
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
	_, err := database.GetDB().Exec("INSERT INTO employees (id, name, age, dept) VALUES (?, ?, ?, ?)", employee.ID, employee.Name, employee.Age, employee.Dept)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	var employee Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	_, err := database.GetDB().Exec("UPDATE employees SET name = ?, age = ?, dept = ? WHERE id = ?", employee.Name, employee.Age, employee.Dept, params["id"])
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Employee not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	employee.ID = params["id"]
	json.NewEncoder(w).Encode(employee)
}
