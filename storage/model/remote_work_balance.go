package model

type RemoteWorkBalance struct {
	ID            int    `json:"id"`
	OrgID         string `json:"org_id"`
	EmployeeID    string `json:"employee_id"`
	Type          string `json:"leave_type"`
	AnnualBalance int    `json:"balance"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
