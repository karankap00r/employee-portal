package model

import "time"

type RemoteWorkRequest struct {
	ID         int       `json:"id"`
	OrgID      int       `json:"org_id"`
	EmployeeID string    `json:"employee_id"`
	Type       string    `json:"type"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	Reason     string    `json:"reason"`
	Status     string    `json:"status"`
	UpdatedBy  string    `json:"updated_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
