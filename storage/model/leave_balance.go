package model

type LeaveBalance struct {
	ID            int    `json:"id"`
	OrgID         string `json:"org_id"`
	EmployeeID    string `json:"employee_id"`
	LeaveType     string `json:"leave_type"`
	AnnualBalance int    `json:"balance"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
