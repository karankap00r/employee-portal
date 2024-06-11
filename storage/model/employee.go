package model

import "time"

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
