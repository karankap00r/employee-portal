package model

import (
	"fmt"
	"time"
)

type Employee struct {
	ID         int       `json:"id"`
	EmployeeID string    `json:"employee_id"`
	Name       string    `json:"name"`
	Position   string    `json:"position"`
	Email      string    `json:"email"`
	Salary     int       `json:"salary"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// String Method to return the string representation of the Employee struct
func (e *Employee) String() string {
	return "Employee ID: " + e.EmployeeID + ", Name: " + e.Name + ", Position: " + e.Position + ", Email: " + e.Email + ", Salary: " + fmt.Sprint(e.Salary)
}
