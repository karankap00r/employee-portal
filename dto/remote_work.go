package dto

type GetRemoteWorkBalanceRequest struct {
	OrgID      int    `json:"org_id"`
	EmployeeID string `json:"employee_id"`
}

type GetRemoteWorkBalanceResponse struct {
	EmployeeID    string `json:"employee_id"`
	Type          string `json:"type"`
	AnnualBalance int    `json:"annual_balance"`
}

type RaiseRemoteWorkRequest struct {
	OrgID      int    `json:"org_id"`
	EmployeeID string `json:"employee_id"`
	Type       string `json:"type"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Reason     string `json:"reason"`
}

type UpdateRemoteWorkRequest struct {
	OrgID     int    `json:"org_id"`
	RequestID int    `json:"request_id"`
	Status    string `json:"status"`
	UpdatedBy string `json:"updated_by"`
}

type GetRemoteWorkRequestsInRangeRequest struct {
	OrgID     int    `json:"org_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type GetRemoteWorkRequestsInRangeResponse struct {
	ID         int    `json:"id"`
	EmployeeID string `json:"employee_id"`
	Type       string `json:"type"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Status     string `json:"status"`
}
