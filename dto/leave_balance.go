package dto

type GetLeaveBalanceRequest struct {
	OrgID      int    `json:"org_id"`
	EmployeeID string `json:"employee_id"`
}

type GetLeaveBalanceResponse struct {
	EmployeeID    string `json:"employee_id"`
	LeaveType     string `json:"leave_type"`
	AnnualBalance int    `json:"annual_balance"`
}

type RaiseLeaveRequest struct {
	OrgID         int    `json:"org_id"`
	EmployeeID    string `json:"employee_id"`
	LeaveCategory string `json:"leave_category"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	Reason        string `json:"reason"`
}

type UpdateLeaveStatusRequest struct {
	OrgID     int    `json:"org_id"`
	LeaveID   int    `json:"leave_id"`
	Status    string `json:"status"`
	UpdatedBy string `json:"updated_by"`
}

type GetLeavesInRangeRequest struct {
	OrgID     int    `json:"org_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type GetLeavesInRangeResponse struct {
	ID            int    `json:"id"`
	EmployeeID    string `json:"employee_id"`
	LeaveCategory string `json:"leave_category"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	Status        string `json:"status"`
}
